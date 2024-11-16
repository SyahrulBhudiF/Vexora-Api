package services

import (
	"context"
	"fmt"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/sirupsen/logrus"
	"net/http"
	"slices"
)

type ImageKitService struct {
	imageKit *imagekit.ImageKit
	ctx      context.Context
}

func NewImageKitService(privateKey string, publicKey string, urlEndpoint string) *ImageKitService {
	imageKit := imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey:  privateKey,
		PublicKey:   publicKey,
		UrlEndpoint: urlEndpoint,
	})

	return &ImageKitService{
		imageKit,
		context.Background(),
	}
}

func (service *ImageKitService) IsValidImage(imageBuff []byte) error {
	mime := http.DetectContentType(imageBuff)

	validImgMimes := []string{
		"image/jpeg",
		"image/png",
		"image/jpg",
	}

	if slices.Contains(validImgMimes, mime) {
		return nil
	}

	return fmt.Errorf("invalid image mime type: %s", mime)
}

func (service *ImageKitService) UploadImage(image string, folderPath string, fileName string) (*uploader.UploadResponse, error) {
	uploadResponse, err := service.imageKit.Uploader.Upload(
		service.ctx,
		image,
		uploader.UploadParam{
			FileName: fileName,
			Folder:   folderPath,
		},
	)

	logrus.Info("upload response: %+v", uploadResponse)

	if err != nil {
		return nil, err
	}

	return uploadResponse, nil
}

func (service *ImageKitService) DeleteImage(imageId string) error {
	_, err := service.imageKit.Media.DeleteFile(service.ctx, imageId)

	if err != nil {
		return err
	}

	return nil
}
