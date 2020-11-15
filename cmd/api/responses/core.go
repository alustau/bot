package responses


const (
	notFoundMessage = "Not Found"
)

//NotFoundMessage is return message default
type NotFoundResponse struct {
	Message string `json:"message"`
}

//NotFoundMessage is return message default
type ErrorResponse struct {
	Message string `json:"error"`
}

//NewNotFoundMessage ...
func NewNotFoundResponse() NotFoundResponse {
	return NotFoundResponse{notFoundMessage}
}

