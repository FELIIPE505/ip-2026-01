/*
 14. Desenvolva um programa para calcular e imprimir o preço final de um carro. O valor do preço inicial de

fábrica é fornecido pelo usuário. O carro pode ter as seguintes opções:
a) Ar condicionado (R$ 1750,00)
b) Pintura metálica (R$ 800,00)
c) Vidro elétrico (R$ 1200,00)
d) Direção hidráulica (R$ 2000,00)
*/
package main

import (
	f "fmt"
)

func main() {
	var pinicial float64
	var escolha string
	f.Println("qual o valor do carro? ")
	f.Scan(&pinicial)
	f.Println("ESCOLHA OPCOES a, b, c, d(sem espaços): ")
	f.Scan(&escolha)

	pfinal := pinicial

	for _, letra := range escolha {
		switch letra {
		case 'a':
			pfinal += 1750
		case 'b':
			pfinal += 800
		case 'c':
			pfinal += 1200
		case 'd':
			pfinal += 2000
		default:
			f.Printf("opção %v inválida!", string(letra))
		}
	}

	f.Printf("O preço final do carro é: R$ %.2f", pfinal)
}
