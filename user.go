package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Username	string
	FullName	string
	Email		string
}

func allUsers(w http.ResponseWriter, r *http.Request)  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	_ = json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("New User Endpoint Hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	username := vars["username"]
	fullname := vars["fullname"]
	email := vars["email"]

	db.Create(&User{Username:username, FullName:fullname, Email:email})
	fmt.Fprintf(w, "New User Successfully Created")

}

func deleteUser(w http.ResponseWriter, r *http.Request)  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	username := vars["username"]

	var user User
	db.Where("username=?", username).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request)  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	username := vars["username"]
	fullname := vars["fullname"]
	email := vars["email"]

	var user User
	db.Where("username=?", username).Find(&user)

	user.FullName = fullname
	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}
