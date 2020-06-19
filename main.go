package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
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
