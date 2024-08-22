package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Servidor rodando na porta 80.")
	http.HandleFunc("/", Hello)
	http.HandleFunc("/ConfigMap", ConfigMap)
	http.HandleFunc("/Secret", Secret)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requisição recebida!")
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Ola, I'm %s. I'm %s", name, age)
	//w.Write([]byte("<h1> ola Mundo FullCycle!!!!"))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, erro := os.ReadFile("/go/myfamily/family.txt")

	if erro != nil {
		log.Fatalf("error reading file: ", erro)
	}

	fmt.Fprintf(w, "My family: %s.", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSOWRD")
	fmt.Fprintf(w, "user: %s. password: %s", user, password)
}
