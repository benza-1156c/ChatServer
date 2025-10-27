package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(id uint64, cooldow time.Duration) (string, error) {
	SECRET := os.Getenv("SECRET_JWT")
	if SECRET == "" {
		return "", jwt.ErrTokenMalformed
	}

	claims := jwt.MapClaims{
		"userId": id,
		"exp":    time.Now().Add(time.Hour * cooldow).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParsedToken(tokenStr string) (jwt.MapClaims, error) {
	secretKey := []byte(os.Getenv("SECRET_JWT"))
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
