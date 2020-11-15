package database

import (
	"testing"
)

func TestConnectDatabaseTest(t *testing.T) {
	db := ConnectDatabaseTest()
	defer db.Close()

	_, err := db.Query("show tables")

	if err != nil {
		t.Error(err)
	}
}

