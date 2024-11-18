package main

import (
	"doodocs-backend/internal/config"
	"doodocs-backend/internal/controller"
	"doodocs-backend/internal/middleware"
	"doodocs-backend/internal/service"
	"log"

	_ "doodocs-backend/internal/docs"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Doodocs Backend API
// @version 1.0
// @description This is a REST API for handling archives and sending emails.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	archiveService := service.NewArchiveService()
	mailService := service.NewMailService(cfg)

	archiveController := controller.NewArchiveController(archiveService)
	mailController := controller.NewMailController(mailService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.RequestLogger())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	api := r.Group("/api")
	{
		archive := api.Group("/archive")
		{
			archive.POST("/information", archiveController.GetArchiveInformation) // Route 1
			archive.POST("/files", archiveController.CreateArchive)               // Route 2
		}
		mail := api.Group("/mail")
		{
			mail.POST("/file", mailController.SendMail) // Route 3
		}
	}

	log.Printf("Starting server on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
