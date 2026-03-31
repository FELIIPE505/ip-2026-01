package main

import (
	f "fmt"
)

func main() {
	var valor float64
	f.Printf("DIGITE UM VALOR DA VENDA: ")
	f.Scan(&valor)

	if valor < 10 && valor >= 0 {
		venda := valor + (valor * 0.7)
		f.Printf("O VALOR DA VENDA É: %.2f", venda)
	} else if valor >= 10 && valor < 30 {
		venda := valor + (valor * 0.5)
		f.Printf("O VALOR DA VENDA É: %.2f", venda)
	} else if valor >= 30 && valor < 50 {
		venda := valor + (valor * 0.4)
		f.Printf("O VALOR DA VENDA É: %.2f", venda)
	} else if valor >= 50 {
		venda := valor + (valor * 0.3)
		f.Printf("O VALOR DA VENDA É: %.2f", venda)
	} else if valor < 0 {
		f.Printf("VALOR INVALIDO")
	}
}
