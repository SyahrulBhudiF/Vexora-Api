package entity

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/google/uuid"
)

type Music struct {
	types.Entity
	HistoryUUID uuid.UUID `gorm:"column:history_uuid;index" json:"history_uuid"`
	ID          string    `json:"id"`
	MusicName   string    `json:"playlist_name"`
	Artist      string    `json:"artist"`
	Path        string    `json:"path"`
	Thumbnail   string    `json:"thumbnail"`
}

func NewMusic(historyUUID uuid.UUID, id string, name string, artist string, path string, thumbnail string) *Music {
	return &Music{
		HistoryUUID: historyUUID,
		ID:          id,
		MusicName:   name,
		Artist:      artist,
		Path:        path,
		Thumbnail:   thumbnail,
	}
}

func (m *Music) TableName() string {
	return "music"
}
