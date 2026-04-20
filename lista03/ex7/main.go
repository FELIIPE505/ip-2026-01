package main

import f "fmt"

func main() {
	var n, soma, somapar, maior, menor int
	var qtdn, qtdpar int
	soma = 0
	qtdn = 0
	primeiro := true
	var continuar int

	for {
		f.Println("Insira um número inteiro:")
		f.Scan(&n)
		soma = soma + n
		qtdn = qtdn + 1

		if primeiro {
			maior = n
			menor = n
			primeiro = false
		} else {
			if n > maior {
				maior = n
			}
			if n < menor {
				menor = n
			}
		}
		if n%2 == 0 {
			somapar = somapar + n
			qtdpar = qtdpar + 1
		}
		f.Println("Deseja continuar inserindo dados?(sim=1)")
		f.Scan(&continuar)

		if continuar != 1 {
			break
		}
	}
	f.Println(" a soma dos números digitados:", soma)
	f.Println(" a quantidade de números digitados:", qtdn)
	f.Println(" a média dos números digitados é:", float64(soma)/float64(qtdn))
	f.Println(" o maior número digitado é:", maior)
	f.Println(" o menor número digitado é:", menor)
	f.Println(" a média dos números pares é:", somapar/qtdpar)
	f.Println("a percentagem dos números ímpares entre todos os números digitados é:", 100*(qtdn-qtdpar)/qtdn, "%")
}
