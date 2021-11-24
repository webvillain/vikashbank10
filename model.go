package model

type User struct {
	Id         int    `json:"id"`
	First_Name string `json:"firstname"`
	Last_Name  string `json:"lastname"`
	Email      string `json:"email"`
}
