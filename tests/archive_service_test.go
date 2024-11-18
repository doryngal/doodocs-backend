package tests

import (
	"doodocs-backend/internal/service"
	"fmt"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"os"
	"testing"
)

func TestAnalyzeArchive(t *testing.T) {
	file, err := os.Open("testdata/sample.zip")
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	archiveService := service.NewArchiveService()
	details, err := archiveService.AnalyzeArchive(file, "sample.zip")

	fmt.Printf("%+v\n", details)
	assert.NoError(t, err)
	assert.Equal(t, "sample.zip", details.Filename)
	assert.Equal(t, float64(7), details.TotalFiles)
}

func TestCreateArchive(t *testing.T) {
	archiveService := service.NewArchiveService()

	multipartFile := &multipart.FileHeader{
		Filename: "testfile.txt",
		Header:   map[string][]string{"Content-Type": {"text/plain"}},
	}

	files := []*multipart.FileHeader{multipartFile}

	archive, err := archiveService.CreateArchive(files)
	assert.Error(t, err) // fails because of missing actual files
	assert.Nil(t, archive)
}
