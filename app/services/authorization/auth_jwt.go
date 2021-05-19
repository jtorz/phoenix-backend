package authorization

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// JWTSvc jwt service is used to generate a jwt string and to decode jwt strings.
//
// The slice of bytes represents the string key user to cipher the data in the jwt.
type JWTSvc []byte

// NewJWT returns a new JWT.
func (jwtSvc JWTSvc) NewJWT(authUser baseservice.JWTData) (string, error) {
	if len(jwtSvc) == 0 {
		panic("empty JWT key")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": authUser.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	return token.SignedString([]byte(jwtSvc))
}

// AuthJWT checks the jwt information.
func (jwtSvc JWTSvc) AuthJWT(c *httphandler.Context) (*baseservice.JWTData, error) {
	tokenString, err := jwtSvc.getBearerToken(c)
	if err != nil {
		return nil, err
	}
	_, claims, err := jwtSvc.parseJWT(tokenString)
	if err != nil {
		return nil, err
	}
	return &baseservice.JWTData{
		ID: claims["userID"].(string),
	}, nil
}

// GetBearerToken returns the bearer jwt from the Authorization header.
func (srv JWTSvc) getBearerToken(c *httphandler.Context) (string, error) {
	tok := c.Request.Header.Get("Authorization")
	if tok == "" {
		return "", fmt.Errorf("empty JWT token: %w", baseerrors.ErrAuth)
	}
	if !strings.HasPrefix(tok, "Bearer ") {
		return "", fmt.Errorf("not a Bearer token. %w", baseerrors.ErrAuth)
	}
	tok = tok[7:]
	return tok, nil
}

func (jwtSvc JWTSvc) parseJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected jwt signing method %v: %w", token.Header["alg"], baseerrors.ErrAuth)
		}
		return jwtSvc, nil
	})

	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, fmt.Errorf("unexpected jwt claims: %w", baseerrors.ErrAuth)
	}
	if !token.Valid {
		return nil, nil, fmt.Errorf("invalid jwt: %w", baseerrors.ErrAuth)
	}
	if err := claims.Valid(); err != nil {
		return nil, nil, err
	}
	return token, claims, nil
}
