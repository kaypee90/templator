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
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	router.GET("/", Ping)
	router.POST("/EmailTemplate/", EmailTemplate)
	port := GetPort()
	log.Info("Server running and listening on http://localhost", port)
	router.Run(port)
}

// GetPort : returns port to be used by server **/
func GetPort() string {
	return ":9898"
}
