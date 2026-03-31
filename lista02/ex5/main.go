package main

import (
	f "fmt"
)

func main() {
	var n int
	f.Println("INSIRA UM NÚMERO INTEIRO: ")
	f.Scan(&n)

	if n%5 == 0 && n > 1 || n < 0 {
		f.Printf("O NÚMERO É DIVISIVEL POR 5")
	} else if n == 0 {
		f.Printf("NÚMERO INVÁLIDO")
	} else {
		f.Printf("O NÚMERO NÃO É DIVISIVEL POR 5")
	}
}
