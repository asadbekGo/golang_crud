package main

import (
	"fmt"
	"log"

	"app/config"
	"app/pkg/db"
	"app/storage"
)

func main() {

	cfg := config.Load()

	db, err := db.ConnectionDB(&cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	u := models.User {
		FirstName: "Jahongir",
		LastName: "Normurodov",
	}

	id, err := storage.Create(db, u)
	if err != nil {
		log.Fatalf("error whiling create user: %v", err)
	}

	fmt.Println(id)

	user, err := storage.GetById(db, id)
	if err != nil {
		log.Fatalf("error whiling get by id user: %v", err)
	}

	users, err := storage.GetList(db)
	if err != nil {
		log.Fatalf("error whiling create user: %v", err)
	}

	fmt.Println(users)
}
