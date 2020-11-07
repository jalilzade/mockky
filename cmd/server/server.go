package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//Server is our main server
type Server struct {
	router     *mux.Router
	httpServer *http.Server
	config     *ServerConfig
}

//NewServer is to create httpserver
func NewServer(config *ServerConfig) (*Server, error) {
	// todo: Create HttppServer

	router := mux.NewRouter()
	registerRouters(router)

	//Check for Int64 Overflow
	if math.MaxInt64/time.Second < config.WriteTimeout {
		return nil, errors.New("WriteTimeout overflows int64")
	}

	if math.MaxInt64/time.Second < config.ReadTimeout {
		return nil, errors.New("ReadTimeout overflows int64")
	}

	server := &Server{
		config: config,
		router: router,
		httpServer: &http.Server{
			Handler:      router,
			Addr:         config.Address,
			WriteTimeout: config.WriteTimeout * time.Second,
			ReadTimeout:  config.ReadTimeout * time.Second,
		},
	}

	return server, nil
}

func (s *Server) start() error {
	// todo: Start and Listen HttpServer
	fmt.Printf("Listening on http://%s \n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) stop() error {
	// todo: Stop HttpServer
	return nil
}
