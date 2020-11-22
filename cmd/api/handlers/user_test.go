package handlers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/cgauge/bot/cmd/api/database"
	"github.com/cgauge/bot/cmd/api/handlers"
	"github.com/cgauge/bot/cmd/api/models"
	"github.com/cgauge/bot/cmd/api/repositories"
	"github.com/cgauge/bot/cmd/api/requests"
	"github.com/cgauge/bot/cmd/api/responses"
	router "github.com/cgauge/bot/cmd/api/routers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	db         *sql.DB
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

func TestRequiredParamsUserCreateHandler(t *testing.T) {
	request := &requests.CreateUserRequest{
		"", "", "",
	}

	bodyBuffer := new(bytes.Buffer)
	json.NewEncoder(bodyBuffer).Encode(request)

	w := httptest.NewRecorder()
	h := &handlers.Handler{DB: db}

	r, err := http.NewRequest("POST", "/users", bodyBuffer)

	router.Router(h).ServeHTTP(w, r)

	if err != nil {
		panic(err)
	}

	body, errBody := ioutil.ReadAll(w.Body)

	if errBody != nil {
		t.Error(errBody)
	}

	response := &responses.InvalidParamResponse{}

	json.Unmarshal(body, response)

	if response.Params[0].Error != "Field name is required." {
		t.Errorf("expected: Field name is required., got: %v\n", response.Params[0].Error)
	}
}

func createUser() (user *models.User) {
	repository.Create("Naruto", "naruto@gmail.com", "123")

	return user
}
