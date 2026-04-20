/* A série de Fetuccine é gerada da seguinte forma: os dois primeiros termos (inteiros) são fornecidos pelo
usuário. A partir daí, os termos são gerados com a soma ou subtração dos dois termos anteriores, ou seja:
Ai = Ai-1 + Ai-2, para i ímpar;
Ai = Ai-1 – Ai-2, para i par.
Crie um programa que imprima os N primeiros termos da série de Fetuccine, assumindo que o usuário digitará
um N>=3. */

package main

import f "fmt"

func main() {
	var n, a, b int
	f.Println("Digite os dois primeiros termos:")
	f.Scan(&a, &b) // faremos para a=1 e b=2
	f.Println("Insira a quantidade de termos da série de Fetuccine que serão impressa:")
	f.Scan(&n)
	if n < 3 {
		f.Println("Número inválido")
		return
	}
	f.Print("1 2 ")
	for i := 1; i <= n; i++ {
		var proximo int
		if i%2 == 0 {
			proximo = b - a
		} else {
			proximo = b + a
		}
		f.Print(proximo, " ")
		a, b = b, proximo
	}
	f.Println()
}
