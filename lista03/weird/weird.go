/*
	Considere um algoritmo que recebe como entrada um inteiro positivo n. Se n for par, o algoritmo o divide por dois, e se n for ímpar, o algoritmo o multiplica por três e adiciona um. O algoritmo repete esse processo até que n seja igual a um. Por exemplo, a sequência para n = 3 é a seguinte:

3 -> 10 -> 5 ->16 -> 8 -> 4 -> 2 -> 1
Sua tarefa é simular a execução do algoritmo para um determinado valor de n.

Entrada
A única linha de entrada contém um inteiro n.

Saída
Imprima uma linha que contenha todos os valores de n durante a execução do algoritmo.

# Restrições

1 <= n <= 10^6
*/
package main

import "fmt"

func main() {
	// Lê o número de entrada
	var n int
	fmt.Scan(&n)

	// Fica repetindo até n chegar em 1
	for {
		// Imprime o valor atual de n (sem pular linha, com espaço)
		fmt.Print(n)

		// Se chegou em 1, para tudo
		if n == 1 {
			break
		}

		// Coloca um espaço antes do próximo número
		fmt.Print(" ")

		// Se n for par, divide por 2
		if n%2 == 0 {
			n = n / 2
		} else {
			// Se n for ímpar, multiplica por 3 e soma 1
			n = n*3 + 1
		}
	}

	// Pula uma linha no final
	fmt.Println()
}
