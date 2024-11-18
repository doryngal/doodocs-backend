package tests

import (
	"doodocs-backend/internal/config"
	"doodocs-backend/internal/service"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestProcessSendMail(t *testing.T) {
	cfg, err := config.LoadConfig("../.env")
	if err != nil {
		t.Fatal(err)
	}

	mailService := service.NewMailService(cfg)

	file, err := os.Open("testdata/contract.pdf")
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	err = mailService.ProcessSendMail(file, "contract.pdf", []string{"oringali.d05@gmail.com"}, "application/pdf")
	assert.NoError(t, err)
}

func TestSendMailToEmails(t *testing.T) {
	cfg := &config.Config{
		SMTPHost: "smtp.example.com",
		SMTPPort: "587",
		SMTPUser: "user@example.com",
		SMTPPass: "password",
	}

	mailService := service.NewMailService(cfg)

	err := mailService.SendMailToEmails([]byte("test data"), "testfile.txt", []string{"oringali.d05@gmail.com"})
	assert.Error(t, err) // Without a real SMTP server, this should fail
}
