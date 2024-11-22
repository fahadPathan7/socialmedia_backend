package batman

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"

	// import static data
	// _ "github.com/haquenafeem/ridealike/batman/statik"
	// "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

// Service interface definition
type Service interface {
	Init() error
	GrpcServer() *grpc.Server
	EnableReverseProxy() error
	EnableTLS(cert string, key string) error
	EnableSwagger(prefix string, jsonFilePath string) error
	EnableReflection()
	GWMux() *runtime.ServeMux
	GrpcDialOptions() []grpc.DialOption
	Run(int)
}

type service struct {
	port              int
	addr              string
	mux               *http.ServeMux
	grpcServer        *grpc.Server
	gmux              *runtime.ServeMux
	grpcDialOptions   []grpc.DialOption
	grpcServerOptions []grpc.ServerOption
	tlsConfig         tlsConfig
	tlsEnabled        bool
}

type tlsConfig struct {
	keyPair  *tls.Certificate
	certPool *x509.CertPool
}

func (s *service) Init() error {
	s.grpcServer = grpc.NewServer()
	s.gmux = runtime.NewServeMux()
	s.mux = http.NewServeMux()
	s.grpcDialOptions = []grpc.DialOption{grpc.WithInsecure()}
	s.tlsEnabled = false
	return nil
}

func (s *service) GrpcServer() *grpc.Server {
	return s.grpcServer
}

func (s *service) GWMux() *runtime.ServeMux {
	return s.gmux
}

func (s *service) GrpcDialOptions() []grpc.DialOption {
	return s.grpcDialOptions
}

func (s *service) EnableTLS(cert string, key string) error {
	var err error
	keyPair, err := tls.X509KeyPair([]byte(cert), []byte(key))
	if err != nil {
		return err
	}

	s.tlsConfig.keyPair = &keyPair

	s.tlsConfig.certPool = x509.NewCertPool()
	ok := s.tlsConfig.certPool.AppendCertsFromPEM([]byte(cert))
	if !ok {
		return errors.New("bad certs")
	}

	// config grpc tls certs
	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: s.addr,
		RootCAs:    s.tlsConfig.certPool,
	})
	s.grpcDialOptions = []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	s.grpcServerOptions = []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(s.tlsConfig.certPool, s.addr))}
	s.grpcServer = grpc.NewServer(s.grpcServerOptions...)
	s.tlsEnabled = true

	return nil
}

func (s *service) EnableSwagger(prefix string, jsonFilePath string) error {
	if s.mux == nil {
		return errors.New("no http mux found: server was not created properly")
	}

	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	staticHandler := http.FileServer(statikFS)
	s.mux.Handle(prefix, http.StripPrefix(prefix, staticHandler))

	s.mux.HandleFunc(fmt.Sprintf("%vswagger.json", prefix), func(w http.ResponseWriter, req *http.Request) {
		source, err := os.Open(jsonFilePath)
		if err != nil {
			return
		}
		defer source.Close()
		io.Copy(w, source)
	})
	return nil
}

func (s *service) EnableReflection() {
	reflection.Register(s.grpcServer)
}

func (s *service) enableGrpcGateway() error {
	s.mux.Handle("/", s.gmux)
	return nil
}

func (s *service) Run(port int) {
	// enable gw mux
	s.enableGrpcGateway()

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("could not listen on port %d: %v", port, err)
	}

	var srv *http.Server

	if s.tlsEnabled {
		srv = &http.Server{
			Addr:    s.addr,
			Handler: grpcHandlerFunc(s.grpcServer, allowCORS(s.mux)),
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{*s.tlsConfig.keyPair},
				NextProtos:   []string{"h2"},
			},
		}
		// start server config

		// list all routes
		listAllRoutes(s.mux)
		// run server
		err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))
		if err != nil {
			log.Fatalf("could not start server: %v", err)
		}
	} else {
		srv = &http.Server{
			Addr:    s.addr,
			Handler: grpcHandlerFunc(s.grpcServer, allowCORS(s.mux)),
		}
		// start server config

		// list all routes
		listAllRoutes(s.mux)
		// run server
		err = srv.Serve(conn)
		if err != nil {
			log.Fatalf("could not start server: %v", err)
		}
	}

}

func (s *service) EnableReverseProxy() error {
	return nil
}

// NewService returns a new Service implementation
func NewService(addr string) Service {
	return &service{
		addr: addr,
	}
}

func listAllRoutes(mux *http.ServeMux) {
	fmt.Println("================= Routes =================")
	v := reflect.ValueOf(mux).Elem()
	fmt.Printf("routes: %v\n", v.FieldByName("m"))
	fmt.Println("==========================================")
}

func ListAllRoutes(mux *http.ServeMux) {
	fmt.Println("================= Routes =================")
	v := reflect.ValueOf(mux).Elem()
	fmt.Printf("routes: %v\n", v.FieldByName("m"))
	fmt.Println("==========================================")
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// AllowCORS okay? no?
func AllowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Printf("preflight request for %s", r.URL.Path)
}
