package jwt

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "OY8cdXmWzn5VPMZfmOODASlAUG6jTSZZ"

type jwtCustomClaims struct {
	UserId   uint64 `json:"userId"`
	Admin    bool   `json:"admin"`
	Account  string `json:"account"`
	NickName string `json:"nickName"`
	jwt.StandardClaims
}

func Create(userId uint64, isAdmin bool, account string, nickName string, issuer string) (string, error) {
	claims := &jwtCustomClaims{
		UserId:   userId,
		Admin:    isAdmin,
		Account:  account,
		NickName: nickName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 2).Unix()),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}

func Parse(tokenString string) (*jwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwtCustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func (this *jwtCustomClaims) Refresh() (string, error) {
	this.ExpiresAt = int64(time.Now().Add(time.Hour * 2).Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, this)
	return token.SignedString([]byte(SecretKey))
}
