/*
	Escreva um programa que receba um número inteiro positivo, verifique e informe se ele é ou não um número

triangular. Obs.: Um número é triangular quando é resultado do produto de três números naturais consecutivos.
Exemplo: 24 = 2 x 3 x 4; 120 = 4 x 5 x 6
*/
package main

import f "fmt"

func main() {
	var n int

	f.Println("Insira um número inteiro:")
	f.Scan(&n)
	triangular := false
	for i := 1; i*(i+1)*(i+2) <= n; i++ {
		if n > 0 {
			if i*(i+1)*(i+2) == n {
				triangular = true
				break
			}
		} else {
			f.Println("número inserido inválido")
		}
	}
	if triangular == true {
		f.Println("O número é triangular")
	} else {
		f.Println("O número não é triangular")
	}
}
