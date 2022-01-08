package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	startMux()
}

func startMux() {
	router := mux.NewRouter()
	router.HandleFunc("/repo", repoHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func repoHandler (writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	name := query.Get("name")
	token := query.Get("token")
	createRepo(name, token)
}