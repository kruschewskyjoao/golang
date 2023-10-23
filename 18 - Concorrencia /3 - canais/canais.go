package main

import (
	"fmt"
	"time"
)

// canal tem duas operações: 1 envia um dado e 2 recebe um dado
// bloqueia a execução do programa
func main() {
	canal := make(chan string)
	go escrever("olá", canal)
	fmt.Println("depois da função escrever")
	//	for {
	//		mensagem, aberto := <-canal
	//		if !aberto {
	//			break
	//		}
	//		fmt.Println(mensagem)
	//	}

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("fim")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
