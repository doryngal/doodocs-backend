package tests

import (
	"doodocs-backend/internal/service"
	"fmt"
	"github.com/stretchr/testify/assert"
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
