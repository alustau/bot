package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfiguration struct {
	drive string
	host string
	port int
	username string
	password string
	name string
}

func Connect(configuration *DatabaseConfiguration) (*sql.DB, error) {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		configuration.username,
		configuration.password,
		configuration.host,
		configuration.port,
		configuration.name,
	)

	db, err := sql.Open(configuration.drive, url)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectDatabaseTest() *sql.DB {
	config := &DatabaseConfiguration{
		drive:    "mysql",
		host:     "127.0.0.1",
		port:     3305,
		username: "root",
		password: "root",
		name:     "cgauge_test",
	}

	db, err := Connect(config)

	if err != nil {
		panic(err)
	}

	return db
}

func ConnectDatabase() (*sql.DB, error) {
	config := &DatabaseConfiguration{
		drive:    "mysql",
		host:     "mysql",
		port:     3306,
		username: "root",
		password: "root",
		name:     "cgauge",
	}
	return Connect(config)
}


