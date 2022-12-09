package main

import (
	// "fmt"
	"log"

	"github.com/asadbekGo/golang_crud/config"
	"github.com/asadbekGo/golang_crud/pkg/db"
	// "github.com/asadbekGo/golang_crud/models"
	"github.com/asadbekGo/golang_crud/storage"
)

func main() {

	cfg := config.Load()

	db, err := db.ConnectionDB(&cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// u := models.User{
	// 	FirstName: "Jahongir",
	// 	LastName:  "Normurodov",
	// }

	// id, err := storage.Create(db, u)
	// if err != nil {
	// 	log.Fatalf("error whiling create user: %v", err)
	// }

	// fmt.Println(id)

	// user, err := storage.GetById(db, id)
	// if err != nil {
	// 	log.Fatalf("error whiling get by id user: %v", err)
	// }

	// users, err := storage.GetList(db)
	// if err != nil {
	// 	log.Fatalf("error whiling create user: %v", err)
	// }



	// {
	// 	user1 := models.User{
	// 		"85da31d9-a309-48f0-9c84-8ad87225db8a",
	// 		"",
	// 		"",
	// 	}

	// 	rowsAffected, err := storage.Update(db, user1)
	// 	if err != nil {
	// 		log.Fatalf("error whiling update user: %v", err)
	// 	}

	// 	if rowsAffected == 0 {
	// 		log.Fatalf("not found rows")
	// 	}

	// 	user, err := storage.GetById(db, user1.Id)
	// 	if err != nil {
	// 		log.Fatalf("error whiling get by id user: %v", err)
	// 	}

	// 	fmt.Println(user)
	// }

	{
		err := storage.Delete(db, "85da31d9-a309-48f0-9c84-8ad87225db8a")
		if err != nil {
			log.Fatalf("error whiling delete user: %v", err)
		}
	}
}
