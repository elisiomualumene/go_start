package main

import (
	database "go_start/src/infra/database/gorm"
	"go_start/src/shared/initializers"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectDB()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Golang Start Server",
			"status":  true,
		})
	})

	if err := router.Run(); err != nil {
		panic(err)
	}

	log.Println("Server is running on port 3000")
}
