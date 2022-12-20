package routes

import (
	"log"
	"net/http"

	"github.com/Api_Go_Rest/controllers"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")            // aqui exibimos todas as personalides.
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")   // aqui exibimos apenas uma das personalidades
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")      // Aqui utilizamos para criar uma nova personalidade.
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete") // Aqui utilizamos para deletar alguma personalidade.
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")        // Aqui utilizamos para editar alguma personalidade.
	log.Fatal(http.ListenAndServe(":8000", r))
}
