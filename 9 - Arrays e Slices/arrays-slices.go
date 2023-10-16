package main

import (
	"fmt"
	"reflect"
)

// array não é muito utilizado nas aplicações. Geralmente se usa mais slice
//sempre tem que declarar o tamanho do array
// não pode existir tipos de dados diferentes dentro do array
func main() {
	fmt.Println("arr e slices")
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}

	array3 := [...]int{1, 2, 3, 4, 5, 6} //os 3 pontos vai fixar o tamanho do array a depender do que vai ser declarado depois
	fmt.Println(array3)

	//slice \/
	//tamanho é flexivel
	//não precisa se preocupar com tamanho
	//é mais utilizado que array!!!
	slice := []int{10, 11, 22, 23, 44, 123, 12}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))  // []int
	fmt.Println(reflect.TypeOf(array3)) // [6]int

	slice = append(slice, 229) //append copia o slice original e adiciona o 229 no final dele
	fmt.Println(slice)

	slice2 := array2[1:3] // pega o array2 do indice 1 ao indice 3 (não inclui o 3).
	fmt.Println(slice2)

	array2[1] = "posição alterada"
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("___________")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6) // estoura a capacidade de 11
	fmt.Println(len(slice3))   // vai mostrar 12 como tamanho
	fmt.Println(cap(slice3))   // vai criar um novo slice com dobro de capacidade. então 24
}
