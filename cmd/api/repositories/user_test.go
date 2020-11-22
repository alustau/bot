package repositories

import (
	"database/sql"
	"fmt"
	"github.com/cgauge/bot/cmd/api/database"
	"github.com/cgauge/bot/cmd/api/models"
	"os"
	"testing"
)

var (
	db         *sql.DB
	repository UserRepository
)

func setup() {
	db = database.ConnectDatabaseTest()

	repository = NewUserRepository(db)
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

func TestGetAllUsersShouldRetrieveMax2(t *testing.T) {
	users, err := repository.GetAll()

	if err != nil {
		t.Error(err)
	}

	if len(users) > 2 {
		t.Errorf("UserRepository.GetAll() is retrieving more than 2 users.")
	}
}

func TestCreateUser(t *testing.T) {
	user, err := repository.Create("User Fake", "use@fake.com", "e778c568-1b00-11eb-adc1-0242ac120002")

	if err != nil {
		t.Error(err)
	}

	if user.ID == 0 {
		fmt.Printf("%+v\n", user)
		t.Errorf("UserRepository.Create() can't create an user in database.")
	}
}

func TestFindUser(t *testing.T) {
	name := "Denis Alustau"

	userFake := createUser(
		name,
		"denisalustau@gmail.com",
		"e778c568-1b00-11eb-adc1-0242ac120002",
	)

	user, err := repository.Find(userFake.ID)

	if err != nil {
		t.Error(err)
	}

	if user.Name != "Denis Alustau" {
		t.Errorf("User.Name is not equal Denis Alustau")
	}
}

func createUser(name, email, slackId string) *models.User {
	user, err := repository.Create(name, email, slackId)

	if err != nil {
		panic(err)
	}

	return user
}
