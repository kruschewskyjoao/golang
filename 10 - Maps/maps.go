package main

import "fmt"

//no map a chave e valor tem que ser do mesmo tipo
func main() {
	fmt.Println("map")
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"]) //Pedro

	teste1 := map[int8]int8{
		1: 1,
		2: 2,
	}
	fmt.Println(teste1)
	fmt.Println(teste1[1]) // 1

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joao",
			"ultimo":   "cerqueira",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "campus 1",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"]) //joao

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "gemeos",
	}
	fmt.Println(usuario2)
}
