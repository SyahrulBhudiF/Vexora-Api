package entity

import "github.com/SyahrulBhudiF/Vexora-Api/internal/types"

type User struct {
	types.Entity
	Username       string `json:"username"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profile_picture"`
}

func NewUser(username string, name string, email string, password string, profilePicture string) *User {
	return &User{
		Username:       username,
		Name:           name,
		Email:          email,
		Password:       password,
		ProfilePicture: profilePicture,
	}
}
func (u *User) TableName() string {
	return "users"
}
