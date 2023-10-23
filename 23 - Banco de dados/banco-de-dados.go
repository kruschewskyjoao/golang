package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		fmt.Println("Dentro do sql open")
		log.Fatal(erro)
	}
	defer db.Close()
	if erro = db.Ping(); erro != nil {
		fmt.Println("Erro dentro do ping")
		log.Fatal(erro)
	}
	fmt.Println("Conexão aberta!")

	linhas, erro := db.Query("SELECT * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()
	fmt.Println(linhas)
}
