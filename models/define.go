package models

// Message .
type Message struct {
	Name   string   `json:"name"`
	Text   string   `json:"text"`
	Image  string   `json:"image"`
	Commit []string `json:"commit"`
	Praise int      `json:"praise"`
}

// DataBase interface .
type DataBase interface {
	Init()
	AuthenticateUser(userID, pwd string) bool
}
