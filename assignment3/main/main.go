package main

import (
	"log"
	"net/http"
	"os"

	"assignment3/controller"
	"assignment3/mapstore"
	"assignment3/router"
)

func main() {
	controller := &controller.CustomerController{
		Store: mapstore.NewMapStore(),
	}
	r := router.InitializeRoutes(controller)
	log.Println("Listening...")
	b := os.Getenv("PORT")
	s := ":" + b
	http.ListenAndServe(s, r)
}
