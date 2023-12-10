package router

import (
	"api/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/client/{id}/", controller.HandleGetListUsers).Methods("GET")
	r.HandleFunc("/client/", controller.HandlePostListUser).Methods("POST")
	r.HandleFunc("/", controller.Index).Methods("GET")
	fmt.Println("Server: http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
