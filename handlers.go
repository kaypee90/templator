package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Ping : Handles GET ping request **/
func Ping(c *gin.Context) {
	log.Info("Pinging service - Status Healthy")
	var message = ResponseMessage{Message: "Healthy"}
	c.JSON(http.StatusOK, message)
}

// EmailTemplate : Handles POST request for generating templates **/
func EmailTemplate(c *gin.Context) {
	var templateDetails TemplateDetails
	c.BindJSON(&templateDetails)

	isSuccessful, template := GenerateEmailTemplate(&templateDetails)
	var message string
	if isSuccessful {
		message = "Template successfully generated!"
		log.Info("Email Template has successfully been generated!!")
	} else {
		message = "Template generation failed!"
		log.Error("Email Template generation process failed!")
	}

	response := ResponseMessage{message, template}
	c.JSON(http.StatusCreated, response)
}
