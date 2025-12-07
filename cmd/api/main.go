package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/Keracode/vidyarthidesk-backend/config"
	"github.com/Keracode/vidyarthidesk-backend/internal/server"
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
	cfg := config.Load()

	var s = server.New(cfg)
	s.RegisterRoutes()

	go gracefulShutDown(s)
	if err := s.App.Listen(":"+cfg.Port, fiber.ListenConfig{EnablePrefork: true}); err != nil {
		log.Fatalf("Server exited with error %v", err)
	}

	log.Println("Graceful shutdown completed")
}
