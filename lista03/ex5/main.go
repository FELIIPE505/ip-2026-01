/* Escreva um programa que receba a idade, a altura e o peso de várias pessoas. Calcule e imprima:
- a quantidade de pessoas com idade superior a 50 anos;
- a média das alturas das pessoas com idade entre 10 e 20 anos;
- a porcentagem de pessoas com peso inferior a 40 quilos entre todas as pessoas analisadas.
Considere que os dados informados são válidos. Pergunte ao usuário se ele deseja continuar digitando dados ou
não (Exemplo: 1 - Sim, Outro valor diferente de 1 - Não). */

package main

import (
	f "fmt"
)

func main() {
	var idade int
	var altura, peso float64
	var countIdade50, countIdade10a20, countPeso40, totalPessoas int
	var somaAlturaIdade10a20 float64
	for {
		f.Println("Digite a idade, altura e peso da pessoa:")
		f.Scan(&idade, &altura, &peso)
		totalPessoas++
		if idade > 50 {
			countIdade50++
		}
		if idade >= 10 && idade <= 20 {
			countIdade10a20++
			somaAlturaIdade10a20 += altura
		}
		if peso < 40 {
			countPeso40++
		}
		f.Println("Deseja continuar digitando dados? (1 - Sim, Outro valor - Não)")
		var resposta int
		f.Scan(&resposta)
		if resposta != 1 {
			break
		}
	}
	var mediaAlturaIdade10a20 float64
	if countIdade10a20 > 0 {
		mediaAlturaIdade10a20 = somaAlturaIdade10a20 / float64(countIdade10a20)
	}
	var porcentagemPeso40 float64
	if totalPessoas > 0 {
		porcentagemPeso40 = (float64(countPeso40) / float64(totalPessoas)) * 100
	}
	f.Printf("Quantidade de pessoas com idade superior a 50 anos: %d\n", countIdade50)
	f.Printf("Média das alturas das pessoas com idade entre 10 e 20 anos: %.2f\n", mediaAlturaIdade10a20)
	f.Printf("Porcentagem de pessoas com peso inferior a 40 quilos: %.2f%%\n", porcentagemPeso40)
}
