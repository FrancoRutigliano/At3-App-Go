package service

import "github.com/golang-jwt/jwt"

type JwtService struct {
	secretKey string
	duration  string
}

func (j *JwtService) New(secretKey string, tokenDuration string) *JwtService {
	return &JwtService{
		secretKey: secretKey,
		duration:  tokenDuration,
	}
}

func (j *JwtService) GenerateToken(uuid string) (string, error) {

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": uuid,
		"exp":    j.duration,
	})

	token, err := t.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
