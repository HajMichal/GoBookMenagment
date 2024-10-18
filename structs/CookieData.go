package structs

import "github.com/golang-jwt/jwt/v5"

type CookieData struct {
	Id 		int 	`json:"id"`
	Name	string 	`json:"name"`
	Email 	string 	`json:"email"`
	Type 	int 	`json:"type"`
	Token 	string	`json:"token"`
	jwt.RegisteredClaims
}
type Cookie struct {
	Token 	string	`json:"token"`
}
