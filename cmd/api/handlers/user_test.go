package handlers_test

import (
	"database/sql"
	"encoding/json"
	"github.com/cgauge/bot/cmd/api/database"
	"github.com/cgauge/bot/cmd/api/handlers"
	"github.com/cgauge/bot/cmd/api/repositories"
	"github.com/cgauge/bot/cmd/api/responses"
	router "github.com/cgauge/bot/cmd/api/routers"
	"github.com/cgauge/bot/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	db *sql.DB
	repository repositories.UserRepository
)

func setup() {
	db = database.ConnectDatabaseTest()
	repository = repositories.NewUserRepository(db)
}

func tearDown() {
	defer db.Close()

	truncate, err := db.Prepare("truncate table users")

	_, err = truncate.Exec()

	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func TestUserListHandler(t *testing.T) {
	createUser()

	response := httptest.NewRecorder()

	h := &handlers.Handler{db}

	request, err := http.NewRequest("GET", "/users", nil)

	if err != nil {
		t.Error(err)
	}

	router.Router(h).ServeHTTP(response, request)

	var userResponse responses.UserResponse

	err = json.NewDecoder(response.Body).Decode(&userResponse)

	if err != nil {
		t.Error(err)
	}

	if userResponse.Users[0].Name != "Naruto" {
		t.Errorf("data.0.name [%s] is different from expected [Naruto]", userResponse.Users[0].Name)
	}
}


func createUser() (user *models.User) {
	user = &models.User{
		Name:      "Naruto",
		Email:     "naruto@gmail.com",
		SlackId:   "123",
	}

	repository.Create(user)

	return user
}

