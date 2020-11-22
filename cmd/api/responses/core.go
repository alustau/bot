package responses

const (
	notFoundMessage = "Not Found"
)

//NotFoundMessage is return message default
type NotFoundResponse struct {
	Message string `json:"message"`
}

//NotFoundMessage is return message default
type DefaultResponse struct {
	Message string `json:"message"`
}

//NotFoundMessage is return message default
type ErrorResponse struct {
	Message string `json:"error"`
}

//NotFoundMessage is return message default
type ParamResponse struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

//NotFoundMessage is return message default
type InvalidParamResponse struct {
	Params []ParamResponse `json:"errors"`
}

//NewNotFoundMessage ...
func NewNotFoundResponse() NotFoundResponse {
	return NotFoundResponse{notFoundMessage}
}
