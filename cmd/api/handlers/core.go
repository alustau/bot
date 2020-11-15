package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cgauge/bot/cmd/api/responses"
	"net/http"
)

//Handler ...
type Handler struct {
	DB *sql.DB
}

//HandleNotFound ...
func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	ResponseJson(w, r, http.StatusNotFound, responses.NewNotFoundResponse())
}

//ResponseJson ...
func ResponseJson(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)


	fmt.Fprintln(w, &buf)
}
