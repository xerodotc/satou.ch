package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var redirectMap = map[string]string{
	"/":   "https://ja.wikipedia.org/wiki/%E7%A0%82%E7%B3%96",
	"/an": "https://anilist.co/character/126754/Satou-Matsuzaka",
	"/vi": "https://vi.satou.ch/",
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.AppEngine = true
	gin.SetMode(gin.ReleaseMode)

	r.GET("/*path", func(c *gin.Context) {
		path := c.Param("path")
		dest, ok := redirectMap[path]
		if !ok {
			c.Data(http.StatusNotFound, "text/plain", []byte("not found"))
			return
		}
		c.Redirect(http.StatusFound, dest)
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Println("Exiting...")
}
