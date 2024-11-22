package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/fahadPathan7/socialmedia_backend/batman"
	accesscontrol "github.com/fahadPathan7/socialmedia_backend/batman/access_control"
	_ "github.com/fahadPathan7/socialmedia_backend/batman/statik"
	"github.com/joho/godotenv"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/fahadPathan7/socialmedia_backend/batman/auth"
	pb "github.com/fahadPathan7/socialmedia_backend/proto/user"
	"github.com/fahadPathan7/socialmedia_backend/user/handler"
	"github.com/fahadPathan7/socialmedia_backend/user/repository"
	"github.com/fahadPathan7/socialmedia_backend/user/setting"
	"github.com/fahadPathan7/socialmedia_backend/user/validation"
	"github.com/rakyll/statik/fs"
	"github.com/soheilhy/cmux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ms struct {
	addr string
	port string
	srv  *grpc.Server
	mux  *http.ServeMux
}

func (s *ms) Run() {
	// Create the main listener.
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		log.Fatal(err)
	}

	// Create a cmux object.
	tcpm := cmux.New(l)

	// Declare the match for different services required.
	httpl := tcpm.Match(cmux.HTTP1Fast())
	grpcl := tcpm.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	http2 := tcpm.Match(cmux.HTTP2())

	// Link the endpoint to the handler function.
	http.Handle("/", batman.AllowCORS(s.mux))

	// Initialize the servers by passing in the custom listeners (sub-listeners).
	go s.ServeGRPC(grpcl)
	go s.ServeHTTP(httpl)
	go s.ServeHTTP(http2)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a C, handle it
			l.Close()
			os.Exit(1)
		}
	}()

	log.Println("GRPC server started.")
	log.Println("HTTP server started.")
	log.Printf("Server listening on port %s \n", s.port)

	// Start cmux serving.
	if err := tcpm.Serve(); !strings.Contains(err.Error(),
		"use of closed network connection") {
		log.Fatal(err)
	}
}

func (s *ms) ServeGRPC(l net.Listener) {
	if err := s.srv.Serve(l); err != nil {
		log.Fatalf("could not start GRPC sever: %v", err)
	}
}

func (s *ms) ServeHTTP(l net.Listener) {
	if err := http.Serve(l, nil); err != nil {
		log.Fatalf("could not start HTTP server: %v", err)
	}
}

func main() {
	// Init services
	err := godotenv.Load()
	if err != nil {
		os.Setenv("ENV", "DEV")
		fmt.Printf("=============> Error reading environment variable : %v\n", err)
	}
	val := os.Getenv("ENV")
	fmt.Printf("=============> Environment set to : %v\n", val)

	ms := ms{
		addr: os.Getenv("API_ADDRESS"),
		port: os.Getenv("API_PORT"),
	}

	// init repo and service
	dbClient, err := getDBClient()
	if err != nil {
		log.Fatalf("cold not get database client: %v", err)
	}
	repo := repository.NewMongoRepository(dbClient)
	validator := validation.NewRequestValidator()
	svcImpl := handler.NewService(repo, auth.New(), *validator)

	ms.srv = grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				// bearertoken.Interceptor( // Bearer token auth intercept
				// 	bearertoken.NewAuthenticator(),
				// 	setting.NotAuthGuardedEndpoints,
				// ),
				accesscontrol.Interceptor(
					auth.New(),
					setting.NotAuthGuardedEndpoints,
					setting.AccessableRoles,
				),
			),
		),
	)

	// Register GRPC
	pb.RegisterUserServiceServer(ms.srv, svcImpl)
	reflection.Register(ms.srv)

	// Register GRPC gateway
	ms.mux = http.NewServeMux()
	gmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	ms.mux.Handle("/", gmux)

	err = pb.RegisterUserServiceHandlerFromEndpoint(context.Background(), gmux, ms.addr, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("could not register grpc gateway: %v", err)
	}

	// Enable swagger
	err = enableSwagger(ms.mux, "/docs/", "./proto/user/user.swagger.json")
	if err != nil {
		log.Fatalf("could not enable swagger: %v", err)
	}

	ms.Run()
}

func getDBClient() (*mongo.Client, error) {
	cs := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cs))
	if err != nil {
		return nil, err
	}
	// defer client.Disconnect(ctx)

	return client, nil
}

func enableSwagger(mux *http.ServeMux, prefix string, jsonFilePath string) error {
	if mux == nil {
		return errors.New("no http mux found: server was not created properly")
	}

	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	saticHandler := http.FileServer(statikFS)
	mux.Handle(prefix, http.StripPrefix(prefix, saticHandler))

	mux.HandleFunc(fmt.Sprintf("%vswagger.json", prefix), func(w http.ResponseWriter, req *http.Request) {
		source, err := os.Open(jsonFilePath)
		if err != nil {
			return
		}
		defer source.Close()
		io.Copy(w, source)
	})
	return nil
}