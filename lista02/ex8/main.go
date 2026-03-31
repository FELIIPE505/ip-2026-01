/*
	Faça um programa que indique se um número inteiro informado pelo usuário está compreendido

entre 20 e 90 ou não. (20 e 90 não estão na faixa de valores).
*/
package main

import (
	f "fmt"
)

func main() {
	var n int
	f.Println("DIGITE UM NUMERO: ")
	f.Scan(&n)

	if n > 20 && n < 90 {
		f.Printf("O NÚMERO %d ESTÁ REPREENDIDO", n)
	} else {
		f.Printf("O NÚMERO %d NÃO ESTÁ REPREENDIDO", n)
	}
}
