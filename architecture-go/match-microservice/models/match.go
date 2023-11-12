package models

type APIMatch struct {
	MatchedUserID string `json:"matchedUserID"`
	WaitFor       int    `json:"waitFor"`
}
