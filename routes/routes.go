package routes

import (
	"log"
	"net/http"

	"github.com/Api_Go_Rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasAsPersonalidades).Methods("Get")        //O .Methods("Get") serve para derfine a requisiçao Get para um determinado recurso, no caso, uma requisição por id.
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get") // Aqui definimos que quando eu acessar a /api/personalidades/{id} informando o id na URL, teremos como resposta apenas a personalidade daquele id.
	log.Fatal(http.ListenAndServe(":8002", r))                                                   //Qualquer problema em nosso servidor será apontado pelo log.Fatal
}
