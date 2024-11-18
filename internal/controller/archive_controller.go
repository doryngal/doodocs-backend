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

// GetArchiveInformation
// @Summary Get archive information
// @Description Accepts an archive file and returns detailed information about its structure
// @Accept  multipart/form-data
// @Produce  json
// @Param   file formData file true "The archive file to be analyzed"
// @Success 200 {object} model.ArchiveDetails
// @Failure 400 {object} string
// @Router /archive/information [post]
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

// CreateArchive
// @Summary Create an archive
// @Description Accepts multiple files and returns a ZIP archive containing them
// @Accept  multipart/form-data
// @Produce  application/zip
// @Param   files[] formData file true "Files to be archived"
// @Success 200 {file} application/zip
// @Failure 400 {object} string
// @Router /archive/files [post]
func (ac *ArchiveController) CreateArchive(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form data"})
		return
	}

	files := form.File["files[]"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no files provided"})
		return
	}

	archive, err := ac.archiveService.CreateArchive(files)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/zip", archive)
}
