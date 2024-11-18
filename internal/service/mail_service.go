package service

import (
	"bytes"
	"doodocs-backend/internal/config"
	"encoding/base64"
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

type MailService struct {
	cfg *config.Config
}

func NewMailService(cfg *config.Config) *MailService {
	return &MailService{cfg: cfg}
}

func (ms *MailService) SendMailToEmails(fileData []byte, fileName string, emails []string) error {
	if len(emails) == 0 {
		return errors.New("email list cannot be empty")
	}

	auth := smtp.PlainAuth("", ms.cfg.SMTPUser, ms.cfg.SMTPPass, ms.cfg.SMTPHost)

	message := ms.createEmailMessage(fileName, emails, fileData)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", ms.cfg.SMTPHost, ms.cfg.SMTPPort), auth, ms.cfg.SMTPUser, emails, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (ms *MailService) createEmailMessage(fileName string, emails []string, fileData []byte) string {
	var msg bytes.Buffer

	recipientList := strings.Join(emails, ", ")

	// Writing headlines
	msg.WriteString("To: undisclosed-recipients:;\r\n")
	msg.WriteString(fmt.Sprintf("Bcc: %s\r\n", recipientList))
	msg.WriteString(fmt.Sprintf("Subject: File Delivery: %s\r\n", fileName))
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: multipart/mixed; boundary=\"boundary\"\r\n\r\n")

	// Writing the body of the message
	msg.WriteString("--boundary\r\n")
	msg.WriteString("Content-Type: text/plain; charset=\"UTF-8\"\r\n")
	msg.WriteString("Content-Transfer-Encoding: 7bit\r\n\r\n")
	msg.WriteString("Please find the attached file.\r\n\r\n")

	// Writing an attachment
	msg.WriteString("--boundary\r\n")
	msg.WriteString(fmt.Sprintf("Content-Type: application/octet-stream\r\n"))
	msg.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", fileName))
	msg.WriteString("Content-Transfer-Encoding: base64\r\n\r\n")

	// Encoding the file in Base64
	encoded := base64.StdEncoding.EncodeToString(fileData)
	msg.WriteString(encoded)
	msg.WriteString("\r\n")
	msg.WriteString("--boundary--")

	return msg.String()
}