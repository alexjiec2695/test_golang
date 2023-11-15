package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"
)

func Generate() (string, error) {
	minutes := os.Getenv("jwtExpirationTimeInMinutes")
	d, err := strconv.Atoi(minutes)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"name":  "",
		"admin": true,
		"exp":   time.Now().Add(time.Duration(d) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", err
	}

	return t, nil
}
