package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"user/config"
	"user/prisma/db"
	"user/routes"

	"github.com/go-openapi/swag"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func init() {
	godotenv.Load()
}

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if config.Port == nil {
		config.Port = swag.String("8080")
	}

	e := echo.New()

	c := jaegertracing.New(e, nil)
	defer c.Close()
	// Log all requests
	e.Use(echomiddleware.Logger())

	// Recover on panic
	e.Use(echomiddleware.Recover())

	client := db.NewClient()
	client.Connect()

	defer client.Disconnect()

	routes.Run(e, client)

	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", *config.Port)))
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}