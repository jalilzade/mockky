package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/thedevsaddam/gojsonq"
)

func (s *Server) registerRouters() {
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	s.router.HandleFunc("/api/{entity}", s.mainEntityHandler)
}

func (s *Server) mainEntityHandler(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	//todo: use path concat
	absPath, err := filepath.Abs(fmt.Sprintf("./cmd/server/%s/%s.json", s.config.DatabaseRoot, parameter["entity"]))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": false, "err": err})
		return
	}

	fmt.Println(absPath)
	jq := gojsonq.New().File(absPath)
	res := jq.Where("price", ">", 10).Get()

	//products[0].name
	//todo: filter objects accordinf to parameters
	json.NewEncoder(w).Encode(res)
}
