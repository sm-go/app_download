package main

import (
	"app-download/config"
	"app-download/ds"
	"app-download/handler"
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
	router := gin.Default()

	ds := ds.NewDataSource()

	h := handler.NewHandler(&handler.HConfig{
		R:  router,
		DS: ds,
	})

	h.Register()
	server := http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%s", config.App.Port),
		Handler:        h.R,
		ReadTimeout:    15 * time.Minute,
		WriteTimeout:   15 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("Server started and listen on port :", config.App.Port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	//block until ctrl+c is come
	<-c

	//close the server
	log.Println("Closing server...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Println("error on closing server")
		log.Println(err.Error())
		os.Exit(1)
	}
	log.Println("Server closed successfully.")
}
