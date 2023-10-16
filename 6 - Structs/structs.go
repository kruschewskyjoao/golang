package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("aquivo structs")
	var u usuario
	fmt.Println(u)
	u.idade = 28
	u.nome = "joao"
	fmt.Println(u)

	endereco1 := endereco{"rua dos bobos", 1}
	usuario2 := usuario{"joao2", 28, endereco1}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "joao"}
	fmt.Println(usuario3)
}
