package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/webvillain/vikashbank10/database"
)

// list users
func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var db = database.ConnectDb()

	users, err := database.GetAllUser(db)
	if err != nil {
		log.Panic("Unable To Geting User List From Database")
	}
	json.NewEncoder(w).Encode(users)
	defer db.Close()

}

// get single user
func SingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var db = database.ConnectDb()
	opt := mux.Vars(r)
	id := opt["id"]
	newid, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	user, err := database.GetSingleUser(db, int(newid))
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(user)
	defer db.Close()

}

// create a new user
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var db = database.ConnectDb()
	opt := mux.Vars(r)
	firstname := opt["firstname"]
	lastname := opt["lastname"]
	email := opt["email"]
	user, err := database.CreateUser(db, firstname, lastname, email)
	if err != nil {
		log.Panic("Unable To Create New User")
	}
	fmt.Println("User Is Created Successfully.")
	json.NewEncoder(w).Encode(user)
	defer db.Close()

}

// update a user into database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var db = database.ConnectDb()
	opt := mux.Vars(r)
	id := opt["id"]
	firstname := opt["firstname"]
	lastname := opt["lastname"]
	email := opt["email"]
	newid, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Unable To Parse Int")
	}
	err = database.UpdateUser(db, int(newid), firstname, lastname, email)
	if err != nil {
		log.Fatal("Unable To Update User.")
	}
	fmt.Println("User Is Updated Successfully .")
	defer db.Close()

}

// deleting a user into database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var db = database.ConnectDb()
	opt := mux.Vars(r)
	id := opt["id"]
	newid, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	err = database.DeleteUser(db, int(newid))
	if err != nil {
		fmt.Println("Unable To Delete User")
	}
	fmt.Println("User Is Deleted Successfully.")
	defer db.Close()

}
