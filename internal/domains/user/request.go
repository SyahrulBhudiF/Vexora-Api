package user

type RegisterRequest struct {
	Username string `json:"username" validate:"required,max=30"`
	Name     string `json:"name" validate:"required,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,max=30"`
	Password string `json:"password" validate:"required,min=8"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type UpdateProfileRequest struct {
	Name           string `json:"name" validate:"max=50"`
	Username       string `json:"username" validate:"max=30"`
	Email          string `json:"email" validate:"email"`
	ProfilePicture string `json:"profile_picture"`
}

type ChangePasswordRequest struct {
	PreviousPassword string `json:"previous_password" validate:"required,min=8"`
	NewPassword      string `json:"new_password" validate:"required,min=8"`
}
