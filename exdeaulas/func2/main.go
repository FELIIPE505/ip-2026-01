/* Escreva uma função em GO que leia 3 números inteiros e retorne o maior deles.*/
package main

import (
	f "fmt"
)

func main() {
	var a, b, c int
	f.Scan(&a, &b, &c)

	maiorn := maior(a, b, c)
	f.Printf("maior é = %.d", maiorn)
}

func maior(a1, b1, c1 int) int {
	numeros := []int{a1, b1, c1}
	maior := numeros[len(numeros)-1]
	return maior
}
