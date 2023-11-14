package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func init() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
}

func main() {
	fmt.Printf("Escutando na porta %d\n", config.Porta)

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
