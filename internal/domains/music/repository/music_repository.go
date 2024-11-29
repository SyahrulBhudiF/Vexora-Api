package repository

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"gorm.io/gorm"
)

type MusicRepository struct {
	types.Repository[entity.Music]
}

func NewMusicRepository(db *gorm.DB) *MusicRepository {
	return &MusicRepository{Repository: types.Repository[entity.Music]{DB: db}}
}
