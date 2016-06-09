package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tiborv/hilbert-geohash/handlers"
)

func main() {
	port := os.Getenv("PORT")

	router := gin.New()
	router.Use(gin.Logger())
	router.Static("/static", "static")
	handlers.Register(router)
	router.LoadHTMLGlob("templates/*")

	if port == "" {
		router.Run(":" + "3000")
	} else {
		router.Run(":" + port)

	}
}
