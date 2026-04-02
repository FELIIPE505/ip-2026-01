package main

import (
	f "fmt"
)

func main() {
	var n int
	f.Printf("DIGITE UM NUMERO DE 3 ALGARISMOS: ")
	f.Scan(&n)

	if n < 1000 && n > 99 {
		soma := n % 10
		f.Printf("O NUMERO É: %d", soma)
	} else {
		f.Print("NUMERO INVALIDO")
	}
}
