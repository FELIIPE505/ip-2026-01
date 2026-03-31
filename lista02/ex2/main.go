package main

import "fmt"

func main() {
	var n int
	fmt.Printf("Digite um numero: ")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Printf("POSITIVO\n")
	} else if n < 0 {
		fmt.Printf("NEGATIVO\n")
	} else {
		fmt.Printf("NULO\n")
	}
}
