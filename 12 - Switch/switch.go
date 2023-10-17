package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda feira"
	case 3:
		return "Terça feira"
	case 4:
		return "Quarta feira"
	case 5:
		return "Quinta feira"
	case 6:
		return "Sexta feira"
	case 7:
		return "Sabado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda feira"
	default:
		return "Dia inválido"
	}
}

func diaDaSemana3(numero int) string {
	var diaCorreto string
	switch {
	case numero == 1:
		diaCorreto = "Domingo"
		// fallthrough se cair no numero == 1 ele retorna "Segunda feira"
	case numero == 2:
		diaCorreto = "Segunda feira"
	default:
		diaCorreto = "Dia inválido"
	}
	return diaCorreto
}

func main() {
	fmt.Println("switch")
	dia := diaDaSemana(3)
	fmt.Println(dia)
	dia1 := diaDaSemana(200)
	fmt.Println(dia1)
	dia2 := diaDaSemana2(2)
	fmt.Println(dia2)
	dia3 := diaDaSemana2(200)
	fmt.Println(dia3)
}
