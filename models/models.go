package models

import "database/sql"

type Book struct {
	ID		   int				`json:"id" gorm:"autoIncrement;unique;primaryKey"`
	Title      string 			`json:"title" gorm:"unique"`
	ISBN       string 			`json:"isbn" gorm:"unique"`
	Author     string 			`json:"author"`
	Publisher  string 			`json:"publisher"`
	Stock      int 				`json:"stock"`
	Available  int 				`json:"available"`
	RemoveInfo sql.NullString 	`json:"removeinfo"`
}

type User struct {
	ID       int	`json:"id" gorm:"autoIncrement;unique;primaryKey"`
	Email	 string `json:"email" gorm:"unique"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Overdue  int 	`json:"overdue"`
	Type     int	`json:"type"`
}
