package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Ping : Handles GET ping request **/
func Ping(w http.ResponseWriter, r *http.Request) {
	log.Info("Pinging service - Status Healthy")
	var message = ResponseMessage{Message: "Healthy"}
	json.NewEncoder(w).Encode(message)
}

// EmailTemplate : Handles POST request for generating templates **/
func EmailTemplate(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var templateDetails TemplateDetails
	json.Unmarshal(reqBody, &templateDetails)

	isSuccessful := GenerateEmailTemplate(&templateDetails)
	var message string
	if isSuccessful {
		message = "Template successfully generated!"
		log.Info("Email Template has successfully been generated!!")
	} else {
		message = "Template generation failed!"
		log.Error("Email Template generation process failed!")
	}
	response := ResponseMessage{message}
	json.NewEncoder(w).Encode(response)
}
