package main

type User struct {
	ID   string
	Name string
}

// users guarda os usuários indexados pelo ID.
// Começa com 10 usuários de mock.
var users = map[string]User{
	"065537ec-2c69-4004-802b-676a6c87a2a2": {ID: "065537ec-2c69-4004-802b-676a6c87a2a2", Name: "Ana Souza"},
	"f32b68d7-ff0c-4087-afde-652da5442c0a": {ID: "f32b68d7-ff0c-4087-afde-652da5442c0a", Name: "Bruno Lima"},
	"8d86c1ca-85fe-4c73-8f8d-c18c74c8c5c9": {ID: "8d86c1ca-85fe-4c73-8f8d-c18c74c8c5c9", Name: "Carla"},
	"d2017057-7d38-4544-93eb-f463b860b901": {ID: "d2017057-7d38-4544-93eb-f463b860b901", Name: "Diego Alves"},
	"8a5f5c1c-a6b7-4315-a873-1a50ee226bea": {ID: "8a5f5c1c-a6b7-4315-a873-1a50ee226bea", Name: "Elisa Rocha"},
	"5df2710f-8ca0-416d-a361-63e771454d98": {ID: "5df2710f-8ca0-416d-a361-63e771454d98", Name: "Felipe"},
	"56f59fdb-8621-44c2-aac2-26c74099a1e1": {ID: "56f59fdb-8621-44c2-aac2-26c74099a1e1", Name: "Gabriela Costa"},
	"69baa323-3b90-4f23-a14e-f5bf718c700a": {ID: "69baa323-3b90-4f23-a14e-f5bf718c700a", Name: "Heitor Martins"},
	"32f5bd09-5aff-447c-94bf-36651036ea33": {ID: "32f5bd09-5aff-447c-94bf-36651036ea33", Name: "Isabela"},
	"5e6f6cdb-8e65-47f7-8ac8-092e7e736e97": {ID: "5e6f6cdb-8e65-47f7-8ac8-092e7e736e97", Name: "João Pereira"},
}
