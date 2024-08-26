package structs

type CookieData struct {
	Id 		int 	`json:"id"`
	Name	string 	`json:"name"`
	Email 	string 	`json:"email"`
	Type 	int 	`json:"type"`
	Token 	string	`json:"token"`
}