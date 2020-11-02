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

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.AppEngine = true
	gin.SetMode(gin.ReleaseMode)

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://ja.wikipedia.org/wiki/%E7%A0%82%E7%B3%96")
	})
	r.GET("/an", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://myanimelist.net/character/134256/Satou_Matsuzaka")
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
