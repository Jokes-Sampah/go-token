package gotoken

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type TokenData struct {
	jwt.StandardClaims
	Id           string `json:"id"`
	Email        string `json:"em"`
	GoogleUserId string `json:"gid"`
	Name         string `json:"nm"`
	PhoneNumber  string `json:"pn"`
}

func ExtractToken(tokenString string, secretKey string) (*TokenData, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &TokenData{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(secretKey), nil
	})

	if token != nil && token.Valid {
		// Token is either expired or not active yet
		// Just pass, let the service decided
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return (*TokenData)(nil), err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			// Just pass, let the service decided
		} else {
			// Couldn't handle this token
			return (*TokenData)(nil), err
		}
	} else {
		return (*TokenData)(nil), errors.New("invalid structure")
	}

	return token.Claims.(*TokenData), nil
}
