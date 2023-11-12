package models

type Message struct {
	User1ID string `json:"user1ID"`
	User2ID string `json:"user2ID"`
	Message string `json:"message"`
	ID      string `json:"id"`
}
