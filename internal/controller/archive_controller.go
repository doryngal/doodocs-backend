package controller

import (
	"doodocs-backend/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArchiveController struct {
	archiveService *service.ArchiveService
}

func NewArchiveController(service *service.ArchiveService) *ArchiveController {
	return &ArchiveController{archiveService: service}
}

func (ac *ArchiveController) GetArchiveInformation(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to retrieve file"})
		return
	}
	defer file.Close()

	result, err := ac.archiveService.AnalyzeArchive(file, header.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read file"})
		return
	}
	c.JSON(http.StatusOK, result)
}
