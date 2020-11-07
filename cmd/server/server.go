package main

import (
	"encoding/json"
	"fmt"
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
func NewServer(config *ServerConfig) *Server {
	// todo: Create HttppServer

	router := mux.NewRouter()
	registerRouters(router)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	server := &Server{
		config: config,
		router: router,
	}

	server.httpServer = &http.Server{
		Handler:      router,
		Addr:         config.Address,
		WriteTimeout: config.WriteTimeout * time.Second,
		ReadTimeout:  config.ReadTimeout * time.Second,
	}

	return server
}

func (s *Server) start() error {
	// todo: Start and Listen HttpServer
	fmt.Printf("Listening on %s \n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) stop() error {
	// todo: Stop HttpServer
	return nil
}
