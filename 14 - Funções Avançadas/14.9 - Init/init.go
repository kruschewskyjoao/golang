package main

import "fmt"

var n int

func init() {
	fmt.Println("função init")
	n = 10
}

func main() {
	fmt.Println("função main")
	fmt.Println(n)
}
