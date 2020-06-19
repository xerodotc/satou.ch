package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://ja.wikipedia.org/wiki/%E7%A0%82%E7%B3%96")
	})
	r.GET("/an", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://myanimelist.net/character/134256/Satou_Matsuzaka")
	})

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}
