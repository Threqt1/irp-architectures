package models

type Match struct {
	MatchedUserID string `json:"matchedUserID"`
	WaitFor       int    `json:"waitFor"`
}
