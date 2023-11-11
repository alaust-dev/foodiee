package models

type User struct {
	Id      string `json:"id,omitempty" `
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
