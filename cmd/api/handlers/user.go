package handlers

import (
	"encoding/json"
	"github.com/cgauge/bot/cmd/api/repositories"
	"github.com/cgauge/bot/cmd/api/requests"
	"github.com/cgauge/bot/cmd/api/responses"
	"log"
	"net/http"
)

func (h *Handler) UsersListHandler(w http.ResponseWriter, r *http.Request) {
	repository := repositories.NewUserRepository(h.DB)

	users, err := repository.GetAll()

	if err != nil {
		log.Println("Error:", err.Error())
		ResponseJson(w, r, 400, responses.ErrorResponse{"Error when tried fetch users"})
		return
	}

	ResponseJson(w, r, 200, responses.UserResponse{users})
	return
}

func (h *Handler) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	userRequest := new(requests.CreateUserRequest)

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userRequest); err != nil {
		log.Println("UserCreateHandler -> Error: ", err)
		ResponseJson(w, r, http.StatusBadRequest, err)
		return
	}

	if !userRequest.IsValid() {
		ResponseJson(w, r, http.StatusBadRequest, userRequest.GetErrors())
		return
	}

	ResponseJson(w, r, http.StatusBadRequest, responses.DefaultResponse{Message: "Success"})
}
