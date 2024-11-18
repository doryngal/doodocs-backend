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

// @Summary Send a file via email
// @Description Sends a file to the provided list of email addresses
// @Accept  multipart/form-data
// @Produce  json
// @Param   file formData file true "File to be sent via email"
// @Param   emails formData string true "Comma-separated list of email addresses"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /mail/file [post]
func (mc *MailController) SendMail(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to retrieve file"})
		return
	}
	defer file.Close()

	emails := ctx.PostForm("emails")
	emailList := strings.Split(emails, ",")

	err = mc.mailService.ProcessSendMail(file, header.Filename, emailList, header.Header.Get("Content-Type"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "email sent successfully"})
}
