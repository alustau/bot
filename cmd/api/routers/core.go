package routers

import (
	"github.com/cgauge/bot/cmd/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

//Router ...
func Router(h *handlers.Handler) *mux.Router {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(handlers.HandleNotFound)

	router.HandleFunc("/users", h.UsersListHandler).Methods("Get")

	return router
}