package infrastructure

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kindai-csg/D-Chat/domain"
)

type TokenJwt struct {
	secret string
	hour   int
}

func NewTokenJwt(secret string, hour int) *TokenJwt {
	tokenJwt := TokenJwt{
		secret: secret,
		hour:   hour,
	}
	return &tokenJwt
}

func (tokenJwt *TokenJwt) Create(data domain.TokenData) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = data.UserId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenJwt.hour)).Unix()
	tokenString, _ := token.SignedString([]byte(tokenJwt.secret))
	return tokenString
}

func (tokenJwt *TokenJwt) Authentication(tokenString string) (domain.TokenData, error) {
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenJwt.secret), nil
	})

	tokenData := domain.TokenData{}

	if err != nil {
		return tokenData, err
	}

	claims := t.Claims.(jwt.MapClaims)
	now := time.Now().Add(time.Hour * 0).Unix()
	if claims["exp"].(float64) < float64(now) {
		return tokenData, errors.New("expired")
	}

	tokenData.UserId = claims["user_id"].(string)

	return tokenData, nil
}
