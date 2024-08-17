package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Servidor rodando na porta 80.")
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requisição recebida!")
	w.Write([]byte("<h1> ola Mundo FullCycle!!!!"))
}
