package models

type JwtToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expires     int    `json:"expires_in"` //срок действия токена( по умолчанию 3600 секунд)
	Scope       string `json:"scope"`      // область действия токена (доступ к объектам и операциям над ними)
	Jti         string `json:"jti"`        //уникальный идентификатор токена
}
