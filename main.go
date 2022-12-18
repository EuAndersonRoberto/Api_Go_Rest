package main

import (
	"fmt"

	"github.com/Api_Go_Rest/routes"

	"github.com/Api_Go_Rest/models"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Mestre Irineu", Historia: "Raimundo Irineu Serra conheceu por meio de indígenas da região a bebida sagrada ayahuasca."},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	fmt.Println("Iniciando o Go rest")
	routes.HandleRequest()
}
