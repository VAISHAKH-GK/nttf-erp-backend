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

	_ "github.com/Keracode/vidyarthidesk-backend/docs"
	_ "github.com/joho/godotenv/autoload"
)

//	@title			VidyarthiDesk API
//	@version		1.0
//	@description	Backend API for VidyarthiDesk ERP Software
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.email	vaishakh@vaishakhgk.com

//	@license.name	GPL-3.0
//	@license.url	https://www.gnu.org/licenses/gpl-3.0.html

//	@host		localhost:9000
//	@BasePath	/api

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

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
