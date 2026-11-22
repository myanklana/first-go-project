package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type server struct {
	endereco string
}

var users = map[string]User{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == "/users" {
		w.WriteHeader(200)
		for id, user := range users {
			fmt.Fprintln(w, id, user.Name)
		}
	}

	if r.Method == "POST" && r.URL.Path == "/users" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		newId := uuid.New().String()
		newUser := User{
			ID:   newId,
			Name: string(body),
		}

		users[newId] = newUser

		w.WriteHeader(http.StatusCreated)

		fmt.Fprintln(w, "usuario criado ")
	}

}

func main() {
	srv := &server{
		endereco: ":8080",
	}

	log.Println("Servidor esta funcionando")
	http.ListenAndServe(srv.endereco, srv)
}
