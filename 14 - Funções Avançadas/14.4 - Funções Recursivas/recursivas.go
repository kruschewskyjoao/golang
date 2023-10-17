package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	// \/ função recursiva
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

// funções recursivas não são recomendadas para numeros grandes
// se por numero grande vai quebrar
func main() {
	fmt.Println("Funções Recursivas")

	// 1 1 2 3 5 8 13

	posicao := uint(15)

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Println(fibonacci(posicao))
}
