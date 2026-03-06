package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lemodoescoding/go-url-shortener/shortener"
	"github.com/lemodoescoding/go-url-shortener/store"
	"net/http"
	"fmt"
)

type UrlCreationRequest struct {
	LongURL string `json:"long_url" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	shortURL := shortener.GenerateShortLink(creationRequest.LongURL, creationRequest.UserID)
	store.SaveURLMapping(shortURL, creationRequest.LongURL, creationRequest.UserID)

	host := "http://localhost:9808/"
	c.JSON(http.StatusOK, gin.H{
		"message": "short url created successfully",
		"short_url": host + shortURL,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl, err := store.RetreiveInitialURL(shortUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	c.Redirect(http.StatusFound, initialUrl)	
}
