package routes

import (
	"log"
	"net/http"

	"github.com/Api_Go_Rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasAsPersonalidades)
	log.Fatal(http.ListenAndServe(":8002", nil)) //Qualquer problema em nosso servidor será apontado pelo log.Fatal
}
