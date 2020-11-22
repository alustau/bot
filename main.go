package main

import (
	"fmt"
	"github.com/cgauge/bot/cmd/api/database"
	"github.com/cgauge/bot/cmd/api/handlers"
	"github.com/cgauge/bot/cmd/api/routers"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	port         = 8181
	readTimeOut  = 10 * time.Second
	writeTimeout = 10 * time.Second
)

func main() {

	log.Println("Starting Connection MySql")

	db, err := database.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	h := &handlers.Handler{DB: db}

	routes := routers.Router(h)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      routes,
		ReadTimeout:  readTimeOut,
		WriteTimeout: writeTimeout,
	}

	log.Println("Starting ListenAndServe", port)

	if err := s.ListenAndServe(); err != nil {
		log.Println("Error in ListenAndServe. Error:", err)
		os.Exit(1)
	}
}
