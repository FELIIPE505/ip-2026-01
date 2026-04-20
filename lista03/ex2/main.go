/* Elabore um programa que apresente os resultados da soma e da média aritmética dos valores pares situados na
faixa numérica de 50 a 70. */

package main

import "fmt"

func main() {
	soma := 0
	count := 0

	for i := 50; i <= 70; i++ {
		if i%2 == 0 {
			soma += i
			count++
		}
	}

	media := float64(soma) / float64(count)

	fmt.Printf("Soma dos valores pares: %d\n", soma)
	fmt.Printf("Média aritmética: %.2f\n", media)
}
