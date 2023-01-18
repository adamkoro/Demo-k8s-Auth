package jwt

import (
	"demo-k8s-auth/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Create jwt token
func Create(username string) (string, error) {
	claims := model.TokenClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(jwt.TimeFunc().Add(24 * 7 * 3600 * 1000)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

// Parse jwt token
func Parse(tokenString string) (*model.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.TokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// Refresh jwt token
func Refresh(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*model.TokenClaims); ok && token.Valid {
		claims.ExpiresAt = jwt.NewNumericDate(jwt.TimeFunc().Add(24 * 7 * 3600 * 1000))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte("secret"))
	}
	return "", err
}

// Validate jwt token
func Validate(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return false
	}
	if _, ok := token.Claims.(*model.TokenClaims); ok && token.Valid {
		return true
	}
	return false
}

// Generate unique uuid
func GenerateID() uint32 {
	id := uuid.New().ID()
	return id
}
