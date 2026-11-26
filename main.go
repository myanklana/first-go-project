package main

import (
	"encoding/json"
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		usersList := []User{}
		for _, user := range users {
			usersList = append(usersList, user)
		}
		json.NewEncoder(w).Encode(usersList)
		return
	}

	if r.Method == "POST" && r.URL.Path == "/users" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		var req UserRequest

		json.NewDecoder(r.Body).Decode(&req)

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
	if r.Method == "GET" {

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) == 3 && parts[1] == "users" && parts[2] != "" {
			id := parts[2]
			user, ok := users[id]
			if !ok {
				http.Error(w, "usuario nao existe", http.StatusNotFound)
				return
			}
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(user)
		}
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
