package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Ping).Methods("GET")
	router.HandleFunc("/EmailTemplate/", EmailTemplate).Methods("POST")
	port := GetPort()
	log.Info("Server running and listening on http://localhost", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// GetPort : returns port to be used by server
func GetPort() string {
	return ":9898"
}
