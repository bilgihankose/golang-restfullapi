package models

type User struct {
	ID        int    `json:"ID"`
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Interest  []Interest
	Profile   string `json:"Profile"`
}
