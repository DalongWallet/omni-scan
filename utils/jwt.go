package utils

import (
	"github.com/DalongWallet/omni-scan/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(conf.AppConfig.JwtSecret)

type Claims struct {
	ApiKey 		string `json:"api_key"`
	jwt.StandardClaims
}

func GenerateToken(apikey string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(20 * time.Second)

	claims := Claims{
		apikey,
		jwt.StandardClaims{
			ExpiresAt: 		expireTime.Unix(),
			Issuer: 		"omni-scan",
		},
	}

	tokenCliams := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenCliams.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

