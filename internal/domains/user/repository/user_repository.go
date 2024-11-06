package repository

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"gorm.io/gorm"
)

type UserRepository struct {
	types.Repository[entity.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Repository: types.Repository[entity.User]{DB: db}}
}
