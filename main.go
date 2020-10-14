package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
)

func init(){
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	router :=  mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/_ping", Ping).Methods("GET")
	port := ":9898"
	log.Info("Server running and listening on port ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
