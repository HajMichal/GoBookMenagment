package main

import (
	"database/sql"
	"time"
)

type Books struct {
	Title      string
	ISBN       string
	Author     string
	Publisher  string
	Stock      int
	Available  int
	RemoveInfo sql.NullString
}

type Users struct {
	ID       string
	Name     string
	Password string
	Overdue  int
	Type     int
}

type Records struct {
	recordID    string
	bookID      string
	userID      string
	IsReturned  bool
	borrowDate  time.Time
	returnDate  sql.NullTime
	deadline    time.Time
	extendTimes int
}

type AddBookStruct struct {
	Title 		string `json:"title"`
	ISBN		string `json:"isbn"`
	Author		string `json:"author"`
	Publisher 	string `json:"publisher"`
	Stock		int	   `json:"stock"`
	Available 	int	   `json:"available`
}