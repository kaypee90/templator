package main

import (
    "encoding/json"
	"net/http"
	
    log "github.com/sirupsen/logrus"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	log.Info("Pinging service - Status Healthy")
	var message = ResponseMessage { Message: "Healthy" }
	json.NewEncoder(w).Encode(message)
}