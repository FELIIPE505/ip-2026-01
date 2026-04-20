package main

import f "fmt"

func main() {
	var n int

	f.Print("Insira a qunatidade de termos da sequência: ")
	f.Scan(&n)

	a := 0
	b := 1

	for i := 1; i <= n; i++ {
		f.Printf("%d - ", a)

		a, b = b, a+b
	}
}
