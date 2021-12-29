package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"assignment3/domain"
)

type CustomerController struct {
	Store domain.CustomerStore
}

func (c *CustomerController) Add(w http.ResponseWriter, r *http.Request) {
	var customer1 domain.Customer
	err1 := json.NewDecoder(r.Body).Decode(&customer1)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err1.Error(), http.StatusBadRequest)
	}
	err := c.Store.Create(customer1)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(customer1.ID + ": exists"))
		return
	}
	w.Write([]byte(customer1.ID + ": Created"))
	fmt.Println(customer1.ID, ":Customer has been created")
}
func (c *CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err1 := c.Store.Update(key, customer)
	if err1 != nil {
		fmt.Println("Error:", err1)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(key + ": does not Exists"))
		return
	}
	w.Write([]byte(key + ": updated"))
	fmt.Println(key, ":Customer Updated")
}
func (c *CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	err := c.Store.Delete(key)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(key + ": doesn't Exists"))
		return
	}
	w.Write([]byte(key + ": deleted"))
	fmt.Println(key, ":Customer Deleted")
}
func (c *CustomerController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	cust, err := c.Store.GetById(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(key + ": does not Exist"))
		fmt.Println("Error:", err)
		return
	}
	converted, err := json.Marshal(cust)
	if err != nil {
		w.Write([]byte("Error"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(converted)
	fmt.Println(cust)
}
func (c *CustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := c.Store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	converted, err := json.Marshal(customers)
	if err != nil {
		w.Write([]byte("Error"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(converted)
	fmt.Println(customers)
}
