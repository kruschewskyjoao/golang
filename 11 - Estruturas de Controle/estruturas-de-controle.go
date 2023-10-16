package main

import "fmt"

func main() {
	fmt.Println("estruturas de controle")

	numero := 10
	if numero > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor que 15")
	}

	// if init. É limitado somente ao escopo do if. Fora do if ela não existe.
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Número menor que zero")
	}
}
