package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Servidor rodando na porta 80.")
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requisição recebida!")
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Ola, I'm %s. I'm %s", name, age)
	//w.Write([]byte("<h1> ola Mundo FullCycle!!!!"))
}
