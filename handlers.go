package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// PingHandler : Handles GET ping request **/
func PingHandler(c *gin.Context) {
	log.Info("Pinging service - Status Healthy")
	var message = ResponseMessage{Message: "Healthy"}
	c.JSON(http.StatusOK, message)
}

// HomeHandler : Handles serving of api doc html file **/
func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "API Documentation",
	})
}

// EmailTemplateHandler : Handles POST request for generating templates **/
func EmailTemplateHandler(c *gin.Context) {
	var templateDetails TemplateDetails
	c.BindJSON(&templateDetails)

	uploader := CloudinaryUploader{}

	isSuccessful, template := GenerateAndUploadEmailTemplate(&templateDetails, uploader)
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
