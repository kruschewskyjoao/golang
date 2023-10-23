package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá usuarios!"))
}

func main() {
	// HTTP é um protocolo de comunicação - base da comunicação web
	// cliente (req) <-> servidor (res)
	// rotas:
	// 				URI - Identificador do recurso
	//				Método - GET, POST, PUT, DELETE
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
