package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

//exemplo: go run main.go ip --host amazon.com.br
func main() {
	fmt.Println("Ponto de partida")
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}

}
