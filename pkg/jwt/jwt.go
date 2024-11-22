package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username string, secretKey string)(string, error){
	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id": id,
		"user_name": username,
		"expiry": time.Now().Add(10 * time.Minute).Unix(),
	})

	var key []byte = []byte(secretKey)
	var tokenStr, err  = token.SignedString(key)
	if err != nil{
		return "",err
	}
	return tokenStr, nil
}

func ValidateToken(tokenStr, secretKey string) (int64, string, int64, error){
	var key []byte = []byte(secretKey)
	var claims jwt.MapClaims = jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return key,nil
	})

	if err != nil{
		return 0, "",0, err
	}

	if !token.Valid{
		return 0, "", 0, errors.New("invalid jwt token")
	}
	return int64(claims["user_id"].(float64)), claims["user_name"].(string), int64(claims["expiry"].(float64)), nil 
}