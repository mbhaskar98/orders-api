package handler

import "net/http"

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create method called!!!\n"))
}
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List method called!!!\n"))
}
func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetByID method called!!!\n"))
}
func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateByID method called!!!\n"))
}
func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteByID method called!!!\n"))
}
