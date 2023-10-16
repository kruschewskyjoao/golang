package main

import "fmt"

func main() {
	// aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// não pode somar int16 com int32
	var numero1 int16 = 10
	var numero2 int32 = 25
	soma1 := numero1 + int16(numero2)
	fmt.Println(soma1)

	//atribuição
	var variavel1 string = "atribuição tem que por o string"
	variavel2 := "inferencia de tipos"
	fmt.Println(variavel1, variavel2)

	//operadores relacionais - sempre retornam um booleano
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//operadores lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(false || true)
	fmt.Println(!true)

	//operadores unitarios
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	numero--
	numero -= 20
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 3
	fmt.Println(numero)

	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto)
}
