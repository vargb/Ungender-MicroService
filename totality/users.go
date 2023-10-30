package main

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Phone   string `json:"phone"`
	Height  string `json:"height"`
	Married bool   `json:"married"`
}
