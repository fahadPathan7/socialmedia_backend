package batman

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/rakyll/statik/fs"
	"github.com/soheilhy/cmux"
)

// BatService declares a service interface
type BatService interface {
	Init()
	RegisterGRPC(GRPCService)
	RegisterHTTP(HTTPService)
	Run()
}

// GRPCService def
type GRPCService interface {
	Register(l net.Listener) error
}

// HTTPService def
type HTTPService interface {
	Register()
}

// NewBatService returns a new service
func NewBatService() BatService {
	return &svc{}
}

type svc struct {
	rootMux cmux.CMux
	httpMux *http.ServeMux
	l       net.Listener
	grpcl   net.Listener
	httpl   net.Listener
	http2   net.Listener
}

func (s *svc) Init() {
	// Create the main listener.
	var err error
	s.l, err = net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal(err)
	}

	// Create a cmux object.
	s.rootMux = cmux.New(s.l)

	// Declare the match for different services required.
	s.httpl = s.rootMux.Match(cmux.HTTP1Fast())
	s.grpcl = s.rootMux.MatchWithWriters(
		cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	s.http2 = s.rootMux.Match(cmux.HTTP2())

	// // Initialize the servers by passing in the custom listeners (sub-listeners).
	// go Reg()(grpcl)
	// go serveHTTP(httpl)
	// go serveHTTP(http2)
}

func (s *svc) RegisterGRPC(grpcSvc GRPCService) {
	log.Println("registering grpc server")
	err := grpcSvc.Register(s.grpcl)
	if err != nil {
		log.Fatalf("could not register grpc service")
	}
	log.Println("finished registering grpc server")

}

func (s *svc) RegisterHTTP(httpSvc HTTPService) {
	log.Println("registering http server")
	httpSvc.Register()
	go s.serveHTTP(s.httpl)
	go s.serveHTTP(s.http2)
	http.Handle("/", s.httpMux)
}

func (s *svc) serveHTTP(l net.Listener) {
	if err := http.Serve(l, nil); err != nil {
		log.Fatalf("could not start HTTP server: %v", err)
	}
}

// Run runs the server
func (s *svc) Run() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a C, handle it
			s.l.Close()
			os.Exit(1)
		}
	}()

	// Link the endpoint to the handler function.
	// http.Handle("/", s.httpMux)
	// batman.ListAllRouts(mux)
	// mux.Handle("/", gmux)
	// err := enableSwagger(s.httpMux, "/docs/", "./proto/profile/profile.swagger.json")
	// if err != nil {
	// 	log.Fatalf("cold not enable swagger: %v", err)
	// }

	log.Println("grpc server started.")
	log.Println("http server started.")
	log.Println("Server listening on pot", 10000)

	// Start cmux serving.
	if err := s.rootMux.Serve(); !strings.Contains(err.Error(),
		"use of closed network connection") {
		log.Fatal(err)
	}
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

// DifferenceInDays calculates difference between two dates in days
func DifferenceInDays(t1 time.Time, t2 time.Time) int {
	// The leap year 2016 had 366 days.
	// t1 := NewDate(2020, 8, 19)
	// t2 := NewDate(2020, 9, 6)
	days := t2.Sub(t1).Hours() / 24
	return int(math.Ceil(days)) // 366
}

// NewDate to time converter
func NewDate(year int, month time.Month, day int, endOfTheDay bool) time.Time {
	if endOfTheDay {
		return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(year, month, day, 23, 59, 59, 0, time.UTC)
}
