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

type LogoutRequest = RefreshTokenRequest

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type UpdateProfileRequest struct {
	Name     string `json:"name" validate:"max=50"`
	Username string `json:"username" validate:"max=30"`
	Email    string `json:"email" validate:"email"`
}

type ChangePasswordRequest struct {
	PreviousPassword string `json:"previous_password" validate:"required,min=8"`
	NewPassword      string `json:"new_password" validate:"required,min=8"`
}

type VerifyEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyOtpRequest struct {
	Email string `json:"email" validate:"required,email"`
	Otp   string `json:"otp" validate:"required"`
}

type ResetPasswordRequest struct {
	VerifyOtpRequest
	NewPassword string `json:"new_password" validate:"required,min=8"`
}
