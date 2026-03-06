package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lemodoescoding/go-url-shortener/handler"
	// "github.com/lemodoescoding/go-url-shortener/shortener"
	"github.com/lemodoescoding/go-url-shortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello to Go URL Shortener!",
		})
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
