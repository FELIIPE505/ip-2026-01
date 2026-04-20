/*
	Escreva um programa que calcule potências. O usuário deve digitar a base e o expoente, e o programa deve

apresentar o resultado (sem usar o comando pow). Assuma que o usuário irá digitar valores positivos.
*/
package main

import (
	f "fmt"
)

func main() {
	var base, expoente, r float64

	f.Println("Insira a base e o expoente:")
	f.Scan(&base, &expoente)

	r = 1

	for i := 0; i < int(expoente); i++ {
		r *= base
	}
	f.Printf("%.4f^%.4f=%.4f", base, expoente, r)
}
