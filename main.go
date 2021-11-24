package main

import (
	"log"
	"net/http"

	_ "github.com/dixonwille/wmenu"
	"github.com/gorilla/mux"
	"github.com/webvillain/vikashbank10/database"
	"github.com/webvillain/vikashbank10/handlers"
)

func main() {
	database.ConnectDb()
	// var opts int64
	// fmt.Println("Welcome To Vsmart Bank By Vikash Parashar .")
	// fmt.Println("Application Is Starting On Port : 8080 on Your Local Machine.")
	// fmt.Println("Ex : `localhost:8080/bank`")
	// fmt.Println("--------------------------------------------")
	// fmt.Println("--------------------------------------------")
	r := mux.NewRouter()
	// fmt.Println("1.Create A New User .")
	// fmt.Println("2.Get User By Id .")
	// fmt.Println("3.Get All Users .")
	// fmt.Println("4.Delete A User .")
	// fmt.Println("5.Update A User")
	// fmt.Scanf("%d", &opts)

	// switch opts {
	// case 1:
	// 	fmt.Println("Creating A User ....")
	r.HandleFunc("/bank/{firstname}/{lastname}/{email}", handlers.CreateNewUser).Methods("POST")
	// case 2:
	// 	fmt.Println("Geting List Of Users From Database ....")
	r.HandleFunc("/bank", handlers.ListUser).Methods("GET")
	// case 3:
	// 	fmt.Println("Geting User By Id ....")
	r.HandleFunc("/bank/{id}", handlers.SingleUser).Methods("GET")

	// case 4:
	// 	fmt.Println("Deleting A User From Database ....")
	r.HandleFunc("/bank/{id}", handlers.DeleteUser)

	// case 5:
	// 	fmt.Println("Updating A User Into Database ....")
	r.HandleFunc("/bank/{id}/{firstname}/{lastname}/{email}", handlers.UpdateUser).Methods("PUT")

	//	}

	log.Fatal(http.ListenAndServe(":8080", r))

}
