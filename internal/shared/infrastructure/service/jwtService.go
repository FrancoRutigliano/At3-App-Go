package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtService struct {
}

func (j *JwtService) New() *JwtService {
	return &JwtService{}
}

func (j *JwtService) GenerateTokenRegister(data map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	for key, value := range data {
		claims[key] = value
	}

	// Cambiamos a HS256 para firmar con una clave secreta en texto plano
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		return "", fmt.Errorf("SECRET_KEY no está configurado")
	}

	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *JwtService) ValidateTokenFromQuery(tokenS string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenS, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (j *JwtService) GetUUIdFromToken(token *jwt.Token) (string, error) {

	if !token.Valid {
		return "", fmt.Errorf("invalid token on uuid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	id, ok := claims["userId"].(string)
	if !ok {
		return "", fmt.Errorf("uuid not found in claims")
	}

	return id, nil
}
