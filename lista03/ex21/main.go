package main

import f "fmt"

func main() {
	var b, n int

	f.Println("Digite o valor da base:")
	f.Scan(&b)
	f.Println("Digite o valor do exxpoente:")
	f.Scan(&n)

	if b < 2 || n <= 1 {
		f.Println("Valor inserido inválido")
		return
	}
	valor := 1

	for i := 0; i < n; i++ {
		valor = valor * b
	}
	f.Printf("%d elevado a %d é: %d", b, n, valor)
}
