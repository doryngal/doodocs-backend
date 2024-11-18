package service

import (
	"archive/zip"
	"bytes"
	"doodocs-backend/internal/model"
	"errors"
	"io"
	"mime/multipart"
)

type ArchiveService struct{}

func NewArchiveService() *ArchiveService {
	return &ArchiveService{}
}

func (s *ArchiveService) AnalyzeArchive(file multipart.File, fileName string) (*model.ArchiveDetails, error) {
	fileSize := calculateFileSize(file)

	zipReader, err := zip.NewReader(file, fileSize)
	if err != nil {
		return nil, errors.New("file is not a valid archive")
	}

	var totalSize int64
	files := make([]model.FileDetails, 0)

	for _, f := range zipReader.File {
		totalSize += int64(f.UncompressedSize64)
		files = append(files, model.FileDetails{
			FilePath: f.Name,
			Size:     float64(f.UncompressedSize64),
			Mimetype: detectMimeType(f.Name),
		})
	}

	return &model.ArchiveDetails{
		Filename:    fileName,
		ArchiveSize: float64(fileSize),
		TotalSize:   float64(totalSize),
		TotalFiles:  float64(len(files)),
		Files:       files,
	}, nil
}

func (s *ArchiveService) CreateArchive(files []*multipart.FileHeader) ([]byte, error) {
	var buffer bytes.Buffer
	zipWriter := zip.NewWriter(&buffer)

	for _, fileHeader := range files {
		if !isValidMimeType(fileHeader.Header.Get("Content-Type")) {
			return nil, errors.New("invalid file type detected")
		}

		file, err := fileHeader.Open()
		if err != nil {
			return nil, errors.New("failed to open file for archiving")
		}
		defer file.Close()

		zipFile, err := zipWriter.Create(fileHeader.Filename)
		if err != nil {
			return nil, errors.New("failed to create zip file entry")
		}

		if _, err := zipFile.Write(buffer.Bytes()); err != nil {
			return nil, errors.New("failed to write file to archive")
		}
	}

	if err := zipWriter.Close(); err != nil {
		return nil, errors.New("failed to close zip writer")
	}

	return buffer.Bytes(), nil
}

func isValidMimeType(mimeType string) bool {
	validMimeTypes := map[string]bool{
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/xml": true,
		"image/jpeg":      true,
		"image/png":       true,
	}
	return validMimeTypes[mimeType]
}

func calculateFileSize(file multipart.File) int64 {
	if seeker, ok := file.(io.Seeker); ok {
		size, _ := seeker.Seek(0, io.SeekEnd)
		seeker.Seek(0, io.SeekStart)
		return size
	}
	return 0
}

func detectMimeType(filename string) string {
	if len(filename) >= 4 && filename[len(filename)-4:] == ".jpg" {
		return "image/jpeg"
	}
	if len(filename) >= 5 && filename[len(filename)-5:] == ".docx" {
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	}
	if len(filename) >= 4 && filename[len(filename)-4:] == ".pdf" {
		return "application/pdf"
	}
	return "unknown"
}
