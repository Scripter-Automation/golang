package Users

import (
	"encoding/json"
	database "main/pkg/Database"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You got me")
}

var (
	user User
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&user)
	addUser(&user)

	println("just addes user to database")
	json.NewEncoder(w).Encode("you just added a user")
}

func addUser(user *User) {

	query := `INSERT INTO users(first_name, last_name, email) VALUES (?,?,?)`

	database.ExecuteQuery(query, user.FirstName, user.LastName, user.Email)
}
