package controller

import (
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleGetListUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["id"]
	fmt.Println(listID)
	list, err := models.GetListUsers(listID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	listJSON, err := json.Marshal(list)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf(`{"list" : %s}`, listJSON)))
}

func HandlePostListUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listID, err := models.AddUserAtList(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	resp := make(map[string]string)

	resp["id"] = listID

	sereliazedResponse, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(sereliazedResponse)
	fmt.Println("User", user)
}
