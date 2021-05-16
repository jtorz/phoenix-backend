package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthService authorization module service
type AuthService struct {
	//dao *DAOAuth
}

// NewAuthService creates a new AuthService.
func NewAuthService() *AuthService {
	return &AuthService{}
}

// GetJWT returns new JWT.
func (srv *AuthService) GetJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	//return token.SignedString(config.GetJWTKey())
	return token.SignedString("config.GetJWTKey()")
}
