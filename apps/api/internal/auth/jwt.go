package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key") 


// Creates a jwt
func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims);

	return token.SignedString(secretKey)
}


func ValidateToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil;
	})

	if err != nil || !token.Valid {
		return "", err;
	}

	claims := token.Claims.(jwt.MapClaims);
	userId := claims["user_id"].(string);

	return userId, nil;
}