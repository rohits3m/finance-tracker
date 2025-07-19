package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtUserClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateJwt(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtUserClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "finance-tracker-go",
		},
	})

	str, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return str, nil
}

func VerifyJwt(tokenStr string) (*JwtUserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtUserClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*JwtUserClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("unknown claims in token")
	}
}
