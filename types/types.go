package types

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	APIKey   *struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		RefreshToken string `json:"refresh_token"`
		Expiry       string `json:"expiry"`
	} `json:"APIKey"`
}
