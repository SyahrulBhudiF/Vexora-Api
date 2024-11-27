package repository

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"gorm.io/gorm"
)

type HistoryRepository struct {
	types.Repository[entity.History]
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{Repository: types.Repository[entity.History]{DB: db}}
}
