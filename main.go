package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	// init db connection
	log.Info("initialising db")
	db := NewInMemoryDB()

	// init handler
	log.Info("initialising handler")
	gameH := NewGameHandler(log, db)

	// init router
	router := NewRouter(gameH)
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	log.Info("starting server")
	httpServer.ListenAndServe()

}
