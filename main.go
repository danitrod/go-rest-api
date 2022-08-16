package main

import (
	"log"

	"github.com/danitrod/go-rest-api/database"
	"github.com/danitrod/go-rest-api/models"
	"github.com/danitrod/go-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Dan", Story: "I am a programmer"},
		{Name: "John", Story: "I am a teacher"},
	}
	database.ConnectToDB()
	log.Println("Server launching")
	routes.HandleRequest()
}
