package gotoken

import (
	"github.com/golang-jwt/jwt"
)

func ExtractToken(tokenString string, secretKey string, claims jwt.Claims) (jwt.Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(secretKey), nil
	})

	if token.Valid {
		// token is valid
		// let the service decided
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return token.Claims, err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			// Just pass, let the service decided
		} else {
			// Couldn't handle this token
			return token.Claims, err
		}
	} else {
		// Couldn't handle this token
		return token.Claims, err
	}

	return token.Claims, nil
}