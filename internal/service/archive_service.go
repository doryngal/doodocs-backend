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
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		return nil, err
	}

	reader, err := zip.NewReader(bytes.NewReader(buffer.Bytes()), int64(buffer.Len()))
	if err != nil {
		return nil, errors.New("file is not a valid archive")
	}

	var totalSize int64
	files := make([]model.FileDetails, 0)

	for _, f := range reader.File {
		totalSize += int64(f.UncompressedSize64)
		files = append(files, model.FileDetails{
			FilePath: f.Name,
			Size:     float64(f.UncompressedSize64),
			Mimetype: detectMimeType(f.Name),
		})
	}

	return &model.ArchiveDetails{
		Filename:    fileName,
		ArchiveSize: float64(buffer.Len()),
		TotalSize:   float64(totalSize),
		TotalFiles:  float64(len(files)),
		Files:       files,
	}, nil
}

func detectMimeType(filename string) string {
	if len(filename) >= 4 && filename[len(filename)-4:] == ".jpg" {
		return "image/jpeg"
	}

	return "unknown"
}
