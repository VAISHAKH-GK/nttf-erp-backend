package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MagnaBit/nttf-erp-backend/internal/server"
	"github.com/gofiber/fiber/v3"

	_ "github.com/joho/godotenv/autoload"
)

func gracefulShutDown(fiberServer *server.WebServer) {
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
}

func main() {
	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	var s = server.New()
	s.RegisterRoutes()

	go gracefulShutDown(s)
	if err := s.App.Listen(":"+port, fiber.ListenConfig{EnablePrefork: true}); err != nil {
		log.Fatalf("Server exited with error %v", err)
	}

	log.Println("Graceful shutdown completed")
}
