package service

import (
	"encoding/json"
	"errors"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"mime/multipart"
	"net/http"
)

type IService interface {
	DetectMood(file *multipart.FileHeader) (*entity.MoodDetectionResponse, error)
}

type Service struct {
	clientURL string
	clientKey string
}

func NewService(url string, key string) *Service {
	return &Service{clientURL: url, clientKey: key}
}

// DetectMood detects the mood based on the provided image file
func (s *Service) DetectMood(file *multipart.FileHeader) (*entity.MoodDetectionResponse, error) {
	if err := helpers.ValidateImageFile(file); err != nil {
		return nil, err
	}

	req, err := helpers.CreateMultipartRequest(s.clientURL, file, s.clientKey)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("mood detection service returned an error")
	}

	var mood entity.MoodDetectionResponse
	if err := json.NewDecoder(resp.Body).Decode(&mood); err != nil {
		return nil, errors.New("failed to decode mood response")
	}

	return &mood, nil
}
