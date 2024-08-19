package models

import "database/sql"

type Book struct {
	Title      string 			`json:"title"`
	ISBN       string 			`json:"isbn"`
	Author     string 			`json:"author"`
	Publisher  string 			`json:"publisher"`
	Stock      int 				`json:"stock"`
	Available  int 				`json:"available"`
	RemoveInfo sql.NullString 	`json:"removeinfo"`
}

type User struct {
	ID       string	`json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Overdue  int 	`json:"overdue"`
	Type     int	`json:"type"`
}
