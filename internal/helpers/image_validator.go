package helpers

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func ValidateImageFile(file *multipart.FileHeader) error {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return errors.New("unsupported file type: only jpg, jpeg, and png are allowed")
	}

	maxSize := int64(10 << 20) // 10 MB
	if file.Size > maxSize {
		return errors.New("file size exceeds the maximum allowed size of 5 MB")
	}

	return nil
}
