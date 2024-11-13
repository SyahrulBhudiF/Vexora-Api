package entity

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/google/uuid"
)

type History struct {
	types.Entity
	UserId       uuid.UUID      `json:"user_uuid" gorm:"type:uuid;foreignKey:UserId;references:UUID;onDelete:cascade"`
	Mood         string         `json:"mood"`
	PlaylistName string         `json:"playlist_name"`
	Path         string         `json:"path"`
	Thumbnail    string         `json:"thumbnail"`
	Music        []entity.Music `gorm:"foreignKey:HistoryId;references:UUID"`
}

func NewHistory(userId uuid.UUID, mood string, playlistName string, path string, thumbnail string) *History {
	return &History{
		UserId:       userId,
		Mood:         mood,
		PlaylistName: playlistName,
		Path:         path,
		Thumbnail:    thumbnail,
	}
}

func (h *History) TableName() string {
	return "history"
}
