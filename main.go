package main

import (
	"fmt"

	"github.com/gabrielnascimento2048/api-rest-go/database"
	"github.com/gabrielnascimento2048/api-rest-go/models"
	"github.com/gabrielnascimento2048/api-rest-go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia2"},
		{Id: 3, Nome: "Nome3", Historia: "Historia3"},
		{Id: 4, Nome: "Nome4", Historia: "Historia4"},
	}
	database.DBConnection()
	fmt.Println("Starting API REST with Go!")
	routes.HandleRequest()
}
