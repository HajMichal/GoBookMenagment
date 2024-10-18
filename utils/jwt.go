package utils

import (
	"errors"
	"os"
	"time"

	"example.com/m/structs"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int, email string, userType int) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{
        "id": id, 
        "email": email, 
        "type": userType,
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

	secretKey := os.Getenv("SECRET")
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }

 return tokenString, nil
}

func VerifyToken(tokenStr string) (*structs.CookieData, error) {
    secretKey := os.Getenv("SECRET")
    token, err := jwt.ParseWithClaims(tokenStr, &structs.CookieData{},func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
     })

     if err != nil {
        return nil, errors.New("INVALID TOKEN")
     }
    
     if !token.Valid {
        return nil, errors.New("INVALID TOKEN")
     }
    
     userData := token.Claims.(*structs.CookieData)

     return userData, nil
}