package controller

import (
	"doodocs-backend/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type MailController struct {
	mailService *service.MailService
}

func NewMailController(mailService *service.MailService) *MailController {
	return &MailController{mailService: mailService}
}

func (mc *MailController) SendMail(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to retrieve file"})
		return
	}
	defer file.Close()

	mineType := header.Header.Get("Content-Type")
	if !isValidMimeTypeForEmail(mineType) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid file type"})
		return
	}

	emails := ctx.PostForm("emails")
	if emails == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email list is required"})
		return
	}
	emailList := strings.Split(emails, ",")

	fileData := make([]byte, header.Size)
	if _, err := file.Read(fileData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read file"})
		return
	}

	err = mc.mailService.SendMailToEmails(fileData, header.Filename, emailList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func isValidMimeTypeForEmail(mimeType string) bool {
	validMimeTypes := map[string]bool{
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/pdf": true,
	}
	return validMimeTypes[mimeType]
}
