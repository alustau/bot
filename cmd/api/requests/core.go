package requests

import "github.com/cgauge/bot/cmd/api/responses"

type RequestContract interface {
	IsValid() bool

	GetErrors() responses.ErrorResponse
}
