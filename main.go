package main

import (
	"github.com/gin-gonic/gin"
	dotenv "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Set datetime for formating for logging **/
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	err := dotenv.Load()
	if err != nil {
		log.Info(".env file wasn't found - Dev")
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/_ping", PingHandler)
	router.GET("/", HomeHandler)
	router.POST("/api/v1/EmailTemplate/", EmailTemplateHandler)
	port := GetPort()
	log.Info("Server running and listening on http://localhost", port)
	router.Run(port)
}

// GetPort : returns port to be used by server **/
func GetPort() string {
	return ":9898"
}
