package main

import (
	"log"
	defaultHTTP "net/http" // defaultHTTP refers the the golang default package library

	"pdfy/internal/domain/service"
	"pdfy/internal/infrastructure/http"
	"pdfy/internal/infrastructure/http/handlers"
)

func main() {

	templateSvc := service.NewTemplateSvc()
	// this object helps in passing all the dependencies to
	// the handlers
	allDeps := &handlers.HandlerDeps{
		TemplateSvc: templateSvc,
	}

	// handlerSet contains the reference to all the handlers objects
	// this will be used to bind the handlers to there respective handlers
	// in handler.NewHandlerContainer
	handlerSet := handlers.NewHandlerContainer(allDeps)

	// instantiating the server object
	server, err := http.NewServer(handlerSet)
	if err != nil {
		log.Fatal("error in starting the http server : ", err.Error())
	}

	// starts the server
	err = defaultHTTP.ListenAndServe(":8080", server.Router)
	if err != nil {
		log.Fatal("error in starting the server ", err.Error())
	}
}
