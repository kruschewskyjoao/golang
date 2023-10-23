package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	TipoDeEndereco2 := enderecos.TipoDeEndereco("Avenida Paulista")
	TipoDeEndereco := enderecos.TipoDeEndereco("x Paulista")
	fmt.Println(TipoDeEndereco2)
	fmt.Println(TipoDeEndereco)
}
