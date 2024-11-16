package _interface

import "github.com/imagekit-developer/imagekit-go/api/uploader"

type IImageKitService interface {
	// IsValidImage checks if the provided image bytes have a valid image MIME type
	// Returns nil if valid, error if invalid
	IsValidImage(imageBuff []byte) error

	// UploadImage uploads an image to ImageKit
	// image: base64 encoded image or URL of image
	// folderPath: destination folder path in ImageKit
	// fileName: name to save the file as
	// Returns upload response and error if any
	UploadImage(image string, folderPath string, fileName string) (*uploader.UploadResponse, error)

	// DeleteImage deletes an image from ImageKit using its ID
	// imageId: ID of the image to delete
	// Returns error if deletion fails
	DeleteImage(imageId string) error
}
