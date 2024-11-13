package entity

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/google/uuid"
)

type Music struct {
	types.Entity
	HistoryId uuid.UUID `json:"history_uuid" gorm:"type:uuid;foreignKey:HistoryId;references:UUID;onDelete:cascade"`
	MusicName string    `json:"music_name"`
	Path      string    `json:"path"`
	Thumbnail string    `json:"thumbnail"`
	Artist    string    `json:"artist"`
}

func NewMusic(historyId uuid.UUID, musicName string, path string, thumbnail string, artist string) *Music {
	return &Music{
		HistoryId: historyId,
		MusicName: musicName,
		Artist:    artist,
		Path:      path,
		Thumbnail: thumbnail,
	}
}

func (m *Music) TableName() string {
	return "music"
}
