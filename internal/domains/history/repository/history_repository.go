package repository

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HistoryRepository struct {
	types.Repository[entity.History]
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{Repository: types.Repository[entity.History]{DB: db}}
}

func (r *HistoryRepository) GetMostFrequentMoodByUserUUID(userUUID uuid.UUID) (string, error) {
	var result struct {
		Mood  string
		Count int64
	}

	err := r.DB.Table("history").
		Select("mood, COUNT(mood) as count").
		Where("user_uuid = ?", userUUID).
		Group("mood").
		Order("count DESC").
		Limit(1).
		Scan(&result).Error

	if err != nil {
		return "", err
	}

	return result.Mood, nil
}
