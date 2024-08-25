package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int, email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{
        "id": id, 
        "email": email, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

	secretKey := os.Getenv("SECRET")
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }

 return tokenString, nil
}

func VerifyToken(tokenStr string) error {
    secretKey := os.Getenv("SECRET")
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
     })
    
     if err != nil {
        return err
     }
    
     if !token.Valid {
        return errors.New("INVALID TOKEN")
     }
    
     return nil
}