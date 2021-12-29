package router

import (
	"github.com/gorilla/mux"

	"assignment3/controller"
)

func InitializeRoutes(h *controller.CustomerController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/customer", h.GetAll).Methods("GET")
	r.HandleFunc("/api/customer/{id}", h.Get).Methods("GET")
	r.HandleFunc("/api/customer", h.Add).Methods("POST")
	r.HandleFunc("/api/customer/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/api/customer/{id}", h.Delete).Methods("DELETE")
	return r
}
