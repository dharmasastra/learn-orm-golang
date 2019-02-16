package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func handleRequests()  {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{username}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{username}/{fullname}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{username}/{fullname}/{email}", newUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func main()  {
	fmt.Println("GO GORM TUTORIAL")

	initialMigration()

	handleRequests()
}
