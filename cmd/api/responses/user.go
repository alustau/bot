package responses

import "github.com/cgauge/bot/models"

//NotFoundMessage is return message default
type UserResponse struct {
	Users []*models.User `json:"data"`
}
