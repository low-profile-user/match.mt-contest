package models

import (
	"api/db"
	"errors"
)

const MAXIMUM_QTY_USERS = 100

type MaximumUsersInAListError struct{}

func (error *MaximumUsersInAListError) Error() string {
	return "[ ERROR ] Essa Lista não pode ter mais usuários."
}

type List struct {
	ID    string `json:"id"`
	Users []User `json:"users"`
}

func AddUserAtList(user *User) (string, error) {
	var listID string

	db := db.ConectToDatabase()

	if db == nil {
		return "", errors.New("não foi possível conectar no banco de dados")
	}

	query := db.QueryRow("SELECT addUser($1::VARCHAR(255), $2::VARCHAR(255), $3::VARCHAR(255), $4::VARCHAR(255), $5::VARCHAR(255), $6::BOOL);", user.ID, user.Name, user.Email, user.Hash, user.Cellphone, user.Status)

	if err := query.Scan(&listID); err != nil {
		return "", err
	}

	defer db.Close()

	return listID, nil
}

func GetListUsers(listID string) ([]User, error) {

	db := db.ConectToDatabase()

	if db == nil {
		return nil, errors.New("não foi possível conectar no banco de dados")
	}

	rows, err := db.Query("SELECT user_id, username, email, user_hash, celphone, user_status FROM Users WHERE list_id = $1;", listID)

	if err != nil {
		return nil, err
	}

	var users []User
	var user *User
	for rows.Next() {
		user = new(User)
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Hash, &user.Cellphone, &user.Status)
		users = append(users, *user)
	}
	defer db.Close()

	return users, nil
}
