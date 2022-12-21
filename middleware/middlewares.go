package middleware

import "net/http"

// Aqui utlizamos está função para que informe a todas as nossas funções que o nosso cabeçalho é um JSON.
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json") //Utilizamos este comando para informar quando houver uma requisição que no nosso cabeçalho será do tipo JSON.
		next.ServeHTTP(w, r)
	})
}
