package model

type Site struct {
	Title       string
	Description string
	Author      string
}

var site = &Site{
	Title:       "My WebSite",
	Description: "My WebSite Description",
	Author:      "Eak Netpanya",
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{Username: "admin", Password: "password"},
}
