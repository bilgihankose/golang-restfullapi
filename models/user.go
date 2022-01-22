package models

type User struct {
	ID        int    `json:"ID"`
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Interests  []Interest
	Profile   string `json:"Profile"`
}
