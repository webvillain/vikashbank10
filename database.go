package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/webvillain/vikashbank10/model"
)

var createTable = `create table if not exists users(id integer primary key , firstname text not null , lastname  text not null ,email text not null );`

func ConnectDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./bank.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()
	fmt.Println("Successfully Connected To Database.")
	stmt, err := db.Prepare(createTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Table Is Created Successfully.")
	return db
}

// creating a new user into db and also printing that user on console
func CreateUser(database *sql.DB, firstname string, lastname string, email string) (*model.User, error) {
	addingUser := `insert into users (firstname , lastname , email)values(?,?,?);`

	stmt, err := database.Prepare(addingUser)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(firstname, lastname, email)
	if err != nil {
		log.Println("Unable To Create User")
	}
	id, _ := res.LastInsertId()
	rowsaffected, _ := res.RowsAffected()
	fmt.Println(id, "Rows Affectede : ", rowsaffected)
	var newuser *model.User
	newuser.Id = int(id)
	newuser.First_Name = firstname
	newuser.Last_Name = lastname
	newuser.Email = email
	return newuser, nil
}

// getting all users with their info which are already into database
func GetAllUser(database *sql.DB) ([]*model.User, error) {
	listUsers := `select * from users ;`
	var user *model.User
	var users []*model.User
	rows, err := database.Query(listUsers)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email)
	}
	users = append(users, user)
	rows.Close()
	return users, nil
}

func GetSingleUser(database *sql.DB, id int) (*model.User, error) {
	singleUser := `select * from users where id = ?`
	var user *model.User
	row, err := database.Query(singleUser, id)
	if err != nil {
		fmt.Println("Unable To Fatch Data From Database .")
	}
	if row.Next() {
		row.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email)
	}
	row.Close()
	return user, nil
}

func DeleteUser(database *sql.DB, id int) error {
	deleteUser := `delete from users where id = ?;`

	stmt, err := database.Prepare(deleteUser)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal("Unable To Delete User From Database")
	}
	fmt.Println("User Is Deleted Successfully .")
	rowsaffected, _ := res.RowsAffected()
	fmt.Println(id, "Rows Affectede : ", rowsaffected)
	return nil

}

func UpdateUser(database *sql.DB, id int, firstnam string, lastname string, email string) error {
	updateUser := `update users set firstname = ?,lastname = ?,email = ? where id = ?;`
	stmt, err := database.Prepare(updateUser)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(firstnam, lastname, email, id)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("User Is Updated Successfully.")
	n, _ := res.RowsAffected()
	fmt.Println("Rows Affected : ", n)
	return nil

}
