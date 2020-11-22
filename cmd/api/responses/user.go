package responses

import "github.com/cgauge/bot/cmd/api/models"

//NotFoundMessage is return message default
type UserResponse struct {
	Users []*models.User `json:"data"`
}
