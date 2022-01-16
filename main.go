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
	router.HandleFunc("/delete", deleteHandler).Methods("DELETE")
	router.HandleFunc("/health", healthHandler).Methods("GET")
	router.HandleFunc("/readiness", readinessHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func repoHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	name := query.Get("name")
	token := query.Get("token")
	createRepo(name, token)
}

func deleteHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	name := query.Get("name")
	token := query.Get("token")
	owner := query.Get("owner")
	deleteRepo(name, token, owner)
}

func healthHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

func readinessHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
