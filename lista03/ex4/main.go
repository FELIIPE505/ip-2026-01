/* Escreva um programa que receba vários números inteiros e verifique se eles são ou não quadrados perfeitos. O
programa deve terminar quando o usuário informar um número menor ou igual a zero. Obs.: Um número é
quadrado perfeito quando tem um número inteiro como raiz quadrada. Não é permitido usar o comando sqrt
em sua solução. */

package main

import (
	f "fmt"
)

func main() {
	var n int
	f.Println("Digite um número inteiro (0 ou negativo para sair):")
	f.Scan(&n)
	for n > 0 {
		raiz := 0
		for raiz*raiz < n {
			raiz++
		}
		if raiz*raiz == n {
			f.Println(n, "é um quadrado perfeito.")
		} else {
			f.Println(n, "não é um quadrado perfeito.")
		}
		f.Println("Digite um número inteiro (0 ou negativo para sair):")
		f.Scan(&n)
	}

}
