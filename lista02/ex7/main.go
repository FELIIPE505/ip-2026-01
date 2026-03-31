/* 7. Escreva um programa que leia três valores inteiros distintos (assuma que o usuário digitará valores
diferentes entre si) e os armazene nas variáveis A, B e C. Em seguida, descubra o menor valor e o
armazene em uma variável denominada MENOR; o maior valor, coloque-o na variável MAIOR e o
valor intermediário, na variável INTER. Imprima os valores em ordem crescente, ou seja, imprima
as variáveis MENOR, INTER e MAIOR, nessa ordem. */

package main

import (
	f "fmt"
	"sort"
)

func main() {
	var a, b, c int
	f.Printf("DIGITE 3 NUMEROS: ")
	f.Scan(&a, &b, &c)

	numeros := []int{a, b, c}

	var MENOR int
	var INTER int
	var MAIOR int
	sort.Ints(numeros)
	MENOR = numeros[0]
	INTER = numeros[1]
	MAIOR = numeros[2]

	f.Printf("MENOR: %d, INTER: %d, MAIOR: %d", MENOR, INTER, MAIOR)
}
