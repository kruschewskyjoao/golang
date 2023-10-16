package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//as funções podem ter mais de um retorno
// a função recebe dois parametros int8 n1 e n2
// vai ter dois retornos int8
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("função f")
	}
	f()
	var x = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	x("função x")
	resultado := x("Texto func")
	fmt.Println(resultado)

	resultadosSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadosSoma, resultadoSubtracao)

	resultadoSomenteSoma, _ := calculosMatematicos(30, 20)
	fmt.Println(resultadoSomenteSoma)

	_, resultadoSomenteSubtracao := calculosMatematicos(20, 10)
	fmt.Println(resultadoSomenteSubtracao)
}
