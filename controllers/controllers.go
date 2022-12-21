package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Api_Go_Rest/database"
	"github.com/Api_Go_Rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade //"p" será uma lista onde contém todas as personalidades vindas do models.Personalidade.
	database.DB.Find(&p)         //Aqui solicitamos ao database que em conexão com o banco de dados, encontre TODAS as personalidades utilizando o Find() para encontra-las, passamos o endereço de memória "&p", onde "p" é um slice de personalidades.
	json.NewEncoder(w).Encode(p) //Aqui solicitamos a visualização do slice de "p" para o json.
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id) //Aqui utlizamos o First para trazer a PRIMEIRA REFERÊNCIA que ele encontrar. Então passamos o endereço de memória "&personalidade" e passar o "id" pois será assim que identificará a personalidade específica.
	json.NewEncoder(w).Encode(personalidade)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade) //Neste caso, para a criação de uma nova personalidade, vamos utilizar o NewDECODER para descodificar a informação em formato JSON que utilizaremos para salvar no banco de dados.
	database.DB.Create(&novaPersonalidade)             // Aqui de fato estamos salvando essas informações no banco de dados, utilizando o endereço de memória "&novaPersonalidade"
	json.NewEncoder(w).Encode(novaPersonalidade)       // Aqui utilizamos o JSON para que possamos visualizar essa nova personalidade após ter sido gravado no banco de dados.
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
