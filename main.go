package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
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
