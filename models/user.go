package models

import "database/sql"

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	SlackId string `json:"slack_id"`
	CreatedAt sql.NullTime`json:"created_at"`
}