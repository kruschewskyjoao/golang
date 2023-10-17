package main

import "fmt"

func recuperarExecucao() {
	// fmt.Println("Tentativa de recuperar execução")
	if r := recover(); r != nil {
		fmt.Println("execução recuperada com sucesso!")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é exatamente 6")
}

func main() {
	// se a média for 6 vai quebrar o programa. o panic mata a execução do programa
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós execução")
}
