package repositories

import (
	"database/sql"
	"github.com/cgauge/bot/cmd/api/models"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db}
}

func (r *UserRepository) GetAll() (users []*models.User, err error) {
	results, err := r.db.Query("select * from users limit 2")

	if err != nil {
		return nil, err
	}

	for results.Next() {
		var user models.User
		err = results.Scan(&user.ID, &user.Name, &user.Email, &user.SlackId, &user.CreatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) Find(id int) (user models.User, err error) {
	sql := "select * from users where id = ?"

	err = r.db.QueryRow(sql, id).Scan(&user.ID, &user.Name, &user.Email, &user.SlackId, &user.CreatedAt)

	return user, nil
}

func (r *UserRepository) Create(name, email, slackId string) (*models.User, error) {
	user := &models.User{
		Name:    name,
		Email:   email,
		SlackId: slackId,
	}

	insert, err := r.db.Prepare("insert into users (name, email, slack_id) values (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	result, err := insert.Exec(user.Name, user.Email, user.SlackId)

	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	user.ID = int(id)

	user.CreatedAt = sql.NullTime{time.Now(), true}

	return user, nil
}
