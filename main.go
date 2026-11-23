package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type server struct {
	endereco string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == "/users" {
		w.WriteHeader(200)
		for _, user := range users {
			fmt.Fprintln(w, user.Name)
		}
		return
	}

	if r.Method == "POST" && r.URL.Path == "/users" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		name := strings.TrimSpace(string(body))
		if name == "" {
			http.Error(w, "nome nao pode ser nulo", http.StatusBadRequest)
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
		return
	}
	http.Error(w, "not found", http.StatusNotFound)

}

func main() {
	srv := &server{
		endereco: ":8080",
	}

	log.Println("Servidor esta funcionando")
	http.ListenAndServe(srv.endereco, srv)
}
