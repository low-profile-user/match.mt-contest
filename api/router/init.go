package router

import (
	"api/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/client/{id}/", controller.HandleGetListUsers).Methods("GET")
	r.HandleFunc("/client/", controller.HandlePostListUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
