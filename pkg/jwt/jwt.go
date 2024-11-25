package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username string, secretKey string)(string, error){
	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id": id,
		"username": username,
		"exp": time.Now().Add(2 * time.Minute).Unix(),
	})

	var key []byte = []byte(secretKey)
	var tokenStr, err  = token.SignedString(key)
	if err != nil{
		return "",err
	}
	return tokenStr, nil
}

func ValidateToken(tokenStr, secretKey string) (int64, string, error){
	var key []byte = []byte(secretKey)
	var claims jwt.MapClaims = jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return key,nil
	})

	if err != nil{
		return 0, "", err
	}

	if !token.Valid{
		return 0, "",  errors.New("invalid jwt token")
	}
	return int64(claims["id"].(float64)), claims["username"].(string), nil 
}

func ValidateTokenWithoutExpiry(tokenStr, secretKey string)(int64, string, error){
	var key []byte = []byte(secretKey)
	var claims jwt.MapClaims = jwt.MapClaims{}

	var(
		token *jwt.Token
		err error
	)
	token, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token)(interface{}, error){
		return key, nil
	}, jwt.WithoutClaimsValidation())

	if err != nil{
		return 0, "",err
	}
	if !token.Valid{
		return 0, "", errors.New("invalid token")
	}
	return int64(claims["id"].(float64)), claims["username"].(string), nil
}