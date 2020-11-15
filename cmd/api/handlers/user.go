package handlers

import (
	"github.com/cgauge/bot/cmd/api/repositories"
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