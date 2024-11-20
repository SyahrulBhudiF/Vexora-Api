package entity

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
)

type User struct {
	types.Entity
	Username       string           `json:"username" gorm:"unique;not null"`
	Name           string           `json:"name"`
	Email          string           `json:"email"`
	Verify         bool             `json:"verify"`
	Password       string           `json:"password"`
	ProfilePicture string           `json:"profile_picture"`
	FileId         string           `json:"file_id"`
	History        []entity.History `gorm:"foreignKey:UserId;references:UUID"`
}

func NewUser(username string, name string, email string, password string, profilePicture string, fileId string) *User {
	return &User{
		Username:       username,
		Name:           name,
		Email:          email,
		Password:       password,
		ProfilePicture: profilePicture,
		FileId:         fileId,
	}
}
func (u *User) TableName() string {
	return "users"
}
