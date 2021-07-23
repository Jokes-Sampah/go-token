package gotoken

import (
	"github.com/golang-jwt/jwt"
)

type TokenData struct {
	jwt.StandardClaims
	Id string `json:"id"`
	Email string `json:"email"`
	GoogleUserId string `json:"googleUserId"`
	Name string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

func ExtractToken(tokenString string, secretKey string) (*TokenData, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &TokenData{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(secretKey), nil
	})

	if token.Valid {
		// token is valid
		// let the service decided
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return token.Claims.(*TokenData), err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			// Just pass, let the service decided
		} else {
			// Couldn't handle this token
			return token.Claims.(*TokenData), err
		}
	} else {
		// Couldn't handle this token
		return token.Claims.(*TokenData), err
	}

	return token.Claims.(*TokenData), nil
}