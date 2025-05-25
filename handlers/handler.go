package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miraklik/url-shortener/shortener"
	"github.com/miraklik/url-shortener/store"
)

func CreateUrl(c *gin.Context) {
	var req struct {
		InitialLink string `json:"init_link"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err.Error()})
		return
	}

	userID := c.Param("userID")

	shortUrl := shortener.GenerateShortLink(req.InitialLink, userID)

	store.SaveURLMapping(shortUrl, req.InitialLink, userID)

	c.JSON(http.StatusOK, gin.H{
		"message":    "short url created successfully",
		"Short Link": shortUrl,
	})
}

func GetShortUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialURL := store.RetrieveInitialUrl(shortUrl)

	c.Redirect(http.StatusMovedPermanently, initialURL)
}
