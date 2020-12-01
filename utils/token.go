package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	JWTSECRETKEY = "SFJISAJFOASJjj"
	JWTEXPIRESAT = 60 * 60 * 24
)

type CustomerClaims struct {
	// UserId int64
	UserName string
	jwt.StandardClaims
}

// 生成token的
func CreateToken(userName string) (string, error) {

	claim := &CustomerClaims{
		// UserId: userId,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(JWTEXPIRESAT) * time.Second).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := t.SignedString([]byte(JWTSECRETKEY))
	if err != nil {
		return "", err
	}
	return token, nil
}

// 解析token的
func ParseToken(tokenString string) (*CustomerClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing mehtod")
		}
		return []byte(JWTSECRETKEY), nil
	})
	if claim, ok := token.Claims.(*CustomerClaims); ok && token.Valid {
		return claim, nil
	} else {
		return nil, err
	}
}
