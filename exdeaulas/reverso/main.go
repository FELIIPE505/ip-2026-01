package main

import (
	"fmt"
	f "fmt"
)

func main() {
	var n int
	f.Print("Digite quantidade de numeros do array: ")
	f.Scan(&n)

	numeros := make([]int, n)
	f.Println("Digite", n, "numeros inteiros")

	for i := 0; i < n; i++ {
		f.Printf("%d° numero: ", i+1)
		f.Scan(&numeros[i])
	}

	fmt.Println("\nExibindo na ordem inversa:")

	//lista de tras pra frente
	for i := n - 1; i >= 0; i-- {
		f.Printf("%d ", numeros[i])
	}
}
