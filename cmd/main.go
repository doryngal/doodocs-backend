package main

import (
	"doodocs-backend/internal/config"
	"doodocs-backend/internal/controller"
	"doodocs-backend/internal/middleware"
	"doodocs-backend/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	archiveService := service.NewArchiveService()

	archiveController := controller.NewArchiveController(archiveService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.RequestLogger())

	api := r.Group("/api")
	{
		archive := api.Group("/archive")
		{
			archive.POST("/information", archiveController.GetArchiveInformation) // Route 1
			archive.POST("/files", archiveController.CreateArchive)               //Route 2
		}
	}

	log.Printf("Starting server on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
