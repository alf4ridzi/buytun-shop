package dto

type AuthTokenResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type RefreshTokenResponse struct {
	Refresh string `json:"refresh"`
}
