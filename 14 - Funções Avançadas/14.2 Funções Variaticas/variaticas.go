package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 200)
	fmt.Println(totalSoma)
	totalSoma1 := soma()
	fmt.Println(totalSoma1) // não quebra. vai ser 0
	escrever("Olá mundo", 10, 2, 3, 4, 5, 6, 7)
}
