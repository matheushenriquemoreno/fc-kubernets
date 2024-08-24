package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var startedat = time.Now()

func main() {
	fmt.Println("Servidor rodando na porta 80.")
	http.HandleFunc("/", Hello)
	http.HandleFunc("/ConfigMap", ConfigMap)
	http.HandleFunc("/Secret", Secret)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requisição recebida!")
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Ola, I'm %s. I'm %s", name, age)
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

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedat)
	// logica adicionada para simular o caso de uma aplicação que demoraria
	// 10 segundos para estár disponivel,
	// configuração do kubbernets que verifica a saude da aplicação = Probes : startup, liveness, readiness
	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}
