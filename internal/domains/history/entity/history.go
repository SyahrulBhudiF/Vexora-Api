package entity

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/google/uuid"
)

type RandomMusic struct {
	ID        string `json:"id"`
	MusicName string `json:"playlist_name"`
	Artist    string `json:"artist"`
	Path      string `json:"path"`
	Thumbnail string `json:"thumbnail"`
}

type PlaylistResponse struct {
	Music []RandomMusic `json:"music"`
}

type MoodResponse struct {
	Mood      string        `json:"detected_mood"`
	Music     []RandomMusic `json:"music"`
	CreatedAt string        `json:"created_at"`
}

type MostMood struct {
	Mood string `json:"mood"`
}

func NewPlaylist(id string, name string, artist string, path string, thumbnail string) *RandomMusic {
	return &RandomMusic{
		ID:        id,
		MusicName: name,
		Artist:    artist,
		Path:      path,
		Thumbnail: thumbnail,
	}
}

func NewPlaylistResponse(playlists []RandomMusic) *PlaylistResponse {
	return &PlaylistResponse{
		Music: playlists,
	}
}

type MoodDetectionResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

type History struct {
	types.Entity
	UserUUID uuid.UUID      `json:"user_uuid"`
	Mood     string         `json:"mood"`
	Music    []entity.Music `gorm:"foreignKey:history_uuid" json:"music"`
}

func NewHistory(userUUID uuid.UUID, mood string) *History {
	return &History{
		UserUUID: userUUID,
		Mood:     mood,
	}
}

func (h *History) TableName() string {
	return "history"
}
