/* Gabriel inventou um código para representar números naturais, usando uma sequência de zeros e uns. Funciona assim, o número natural é representado pela quantidade de vezes que o padrão "100" aparece na sequência.

Por exemplo, na sequência 11101001010011110, o padrão aparece duas vezes e na sequência 11101010111110111010101 ele não aparece nenhuma vez. Você deve ajudar Gabriel implementar um programa que, dada a sequência de zeros e uns, calcule quantas vezes o padrão "100" aparece nela.

Entrada
A primeira linha da entrada contém um inteiro N, o tamanho da sequência. A segunda linha contém a sequência de N zeros e uns, separados por espaço em branco.

Saída
Seu programa deve imprimir um inteiro, quantas vezes o padrão "100" aparece na sequência.

Restrições
1 ≤ N ≤ 10^4
1≤ N ≤ 10^4 */

package main

import (
	f "fmt"
)

func main() {
	// Lê o tamanho da sequência
	var n int
	f.Scan(&n)

	// Lê todos os números da sequência em uma lista
	sequencia := make([]int, n)
	for i := 0; i < n; i++ {
		f.Scan(&sequencia[i])
	}

	// Conta quantas vezes o padrão 1, 0, 0 aparece
	contador := 0

	// Percorre a sequência, mas para antes dos últimos 2 elementos
	// porque o padrão tem 3 números (precisamos de i, i+1 e i+2)
	for i := 0; i < n-2; i++ {
		if sequencia[i] == 1 && sequencia[i+1] == 0 && sequencia[i+2] == 0 {
			contador++
		}
	}

	f.Println(contador)
}
