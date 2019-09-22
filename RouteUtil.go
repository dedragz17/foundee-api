package FoundeeAPI

import (
	"github.com/foundee-org/foundee-api/controller"
	"github.com/gorilla/mux"
	"log"
)

func RouteUtil(router *mux.Router) {

	log.Print("Adding routes to router...")

	// error handlers
	router.MethodNotAllowedHandler = nil
	router.NotFoundHandler = nil

	//routes
	router.HandleFunc("/", controller.IndexController)
}
