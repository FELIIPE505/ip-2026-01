package main

import f "fmt"

func main() {
	var soma int
	soma = 0
	for i := 1; i <= 20; i++ {
		f.Print(i, ", ")
		soma = soma + i
	}
	f.Println("Soma=", soma)
}
