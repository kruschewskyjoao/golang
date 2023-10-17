package main

import "fmt"

func funcao1() {
	fmt.Println("Funcao 1")
}

func funcao2() {
	fmt.Println("Funcao 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado") // vai ser retornado antes de dar o return na função
	fmt.Println("entrando na função de aprovado ou não")
	media := (n1 + n2) / 2
	if media >= 7 {
		return true
	}
	return false
}

// DEFER = ADIAR
// muito utilizado em banco de dados.
func main() {

	defer funcao1() // vai executar a função 2 primeiro que a função que está com defer
	funcao2()
	fmt.Println(alunoAprovado(7, 8)) //funcao1 vai ser executado dps dessa também
}
