package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRENCIA != PARALELISMO
	go escrever("olá mundo")
	escrever("Programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
