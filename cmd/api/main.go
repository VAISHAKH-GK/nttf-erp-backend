package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MagnaBit/nttf-erp-backend/internal/server"
	"github.com/gofiber/fiber/v3"

	_ "github.com/joho/godotenv/autoload"
)

func gracefulShutDown(fiberServer *server.WebServer, done chan bool) {
	var ctx, stop = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("Shutting down gracefully, Press Ctrl+C for force shutdown")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := fiberServer.Shutdown(ctx); err != nil {
		log.Printf("Server stopped with error %v", err)
	}

	fmt.Println("Stopping Exiting")
	done <- true
}

func main() {
	var port string
	var done = make(chan bool)

	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	var s = server.New()
	s.RegisterRoutes()

	go s.App.Listen(":"+port, fiber.ListenConfig{EnablePrefork: true})
	go gracefulShutDown(s, done)

	<-done
	log.Println("Graceful shutdown completed")
}
