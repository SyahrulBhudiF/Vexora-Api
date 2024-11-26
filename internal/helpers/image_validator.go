package helpers

import (
	"errors"
	"mime/multipart"
)

func ValidateImageFile(file *multipart.FileHeader) error {
	allowedTypes := []string{"image/jpeg", "image/png"}
	fileType := file.Header.Get("Content-Type")

	isValidType := false
	for _, allowedType := range allowedTypes {
		if fileType == allowedType {
			isValidType = true
			break
		}
	}

	if !isValidType {
		return errors.New("file must be an image (JPEG or PNG)")
	}

	return nil
}
