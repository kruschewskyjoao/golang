package main

import (
	"errors"
	"fmt"
)

func main() {
	//4 numeros inteiros no go. Diferentes na quantidades de bits ex int8, int16, int32, int64
	var numero int8 = 100
	numero2 := 1000000000
	fmt.Println(numero)
	fmt.Println(numero2)

	//uint unsyngned int
	var numero3 uint32 = 10000000
	fmt.Println(numero3)

	//alias sao 'apelidos'
	// int32 = RUNE
	var numero4 rune = 123456
	fmt.Println(numero4)

	// byte = UINT8
	var numero5 byte = 123
	fmt.Println(numero5)

	//numeros quebrados float32 e float64
	var numeroReal1 float32 = 123.54
	var numeroReal2 float64 = 123123123123.321
	fmt.Println(numeroReal1, numeroReal2)
	numeroReal3 := 12345.43
	fmt.Println(numeroReal3)
	//FIM numeros reais

	var str string = "texto"
	fmt.Println(str)
	str2 := "texto2"
	fmt.Println(str2)
	//fim string

	//valor zero. Valor inicial de variavel ex: inicializar uma string vazia ou int vazio vai ser zero
	var texto string
	fmt.Println(texto)
	var numeroVazio int
	fmt.Println(numeroVazio)

	//booleano
	var booleano1 bool = true
	fmt.Println(booleano1)
	var booleano2 bool
	fmt.Println(booleano2)

	//erro <nil>
	// para criar um erro tem que usar um pacote errorS
	var erro error
	fmt.Println(erro)
	var erro1 error = errors.New("Erro criado")
	fmt.Println(erro1)
}
