package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/miraklik/url-shortener/handlers"
	"github.com/miraklik/url-shortener/store"
)

func main() {
	router := gin.Default()

	router.POST("/create-short-url/:userID", handlers.CreateUrl)
	router.GET("/:shortUrl", handlers.GetShortUrl)

	store.NewStore()

	if err := router.Run(":8080"); err != nil {
		log.Panicf("Failed to start server: %v", err)
	}
}
