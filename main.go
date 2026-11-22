package main

import (
	"log"
	"net/http"
)

type server struct {
	endereco string
}

var users = []string{"Matheus", "Miguel", "Gi"}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("Hello from server"))

}

func main() {
	srv := &server{
		endereco: ":8080",
	}

	log.Println("Servidor esta funcionando")
	http.ListenAndServe(srv.endereco, srv)
}
