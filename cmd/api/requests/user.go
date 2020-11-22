package requests

import (
	"github.com/cgauge/bot/cmd/api/responses"
	"github.com/cgauge/bot/cmd/api/validators"
)

type CreateUserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	SlackId string `json:"slackId"`
}

func (r *CreateUserRequest) IsValid() bool {
	return r.requiredParams() && validators.IsEmailValid(r.Email)
}

func (r *CreateUserRequest) GetErrors() responses.InvalidParamResponse {
	var params []responses.ParamResponse

	if r.Name == "" {
		params = append(params, responses.ParamResponse{Field: "name", Error: "Field name is required."})
	}

	if r.Email == "" {
		params = append(params, responses.ParamResponse{Field: "email", Error: "Field email is required."})
	}

	if r.Email != "" && !validators.IsEmailValid(r.Email) {
		params = append(params, responses.ParamResponse{Field: "email", Error: "Field email is not a valid email."})
	}

	return responses.InvalidParamResponse{Params: params}
}

func (r *CreateUserRequest) requiredParams() bool {
	return r.Name != "" && r.Email != "" && r.SlackId != ""
}
