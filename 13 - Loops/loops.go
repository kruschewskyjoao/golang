package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	//	i++
	//	fmt.Println("increment")
	//	time.Sleep(time.Second)
	// }
	//fmt.Println(i)

	//for j := 0; j < 10; j += 2 {
	//	fmt.Println("increment j", j)
	//	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Joao", "Davi", "Lucas"}

	// primeiro sempre indice e o segundo é o valor dentro do array
	for indice, nome := range nomes {
		fmt.Println(indice, nome) // 0 João 1 Davi 2 Lucas
	}
	// caso queira só o valor dentro do array
	for _, nome := range nomes {
		fmt.Println(nome) // João Davi Lucas
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// for {
	//	fmt.Println("Executa infinitamente")
	//	time.Sleep(time.Second)
	// }
}
