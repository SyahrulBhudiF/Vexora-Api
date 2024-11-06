package entity

type Token struct {
	RefreshToken string `json:"refresh_token"`
	Username     string `json:"username"`
}
