package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 2)
	}("param")
	fmt.Println(retorno)

	func(texto string) {
		fmt.Println(texto)
	}("olá função anonima")
}
