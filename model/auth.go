package model

type AuthMessage struct {
	Type        string `json:"type"`
	AccessToken string `json:"access_token"`
}
