package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/thedevsaddam/gojsonq"
)

func registerRouters(router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	router.HandleFunc("/api/{entity}", mainEntityHandler)
}

func mainEntityHandler(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	//todo: use path concat
	absPath, err := filepath.Abs(fmt.Sprintf("./cmd/server/database/%s.json", parameter["entity"]))
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
