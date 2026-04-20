package main

import f "fmt"

func main() {
	const numNotas int = 3 // quantidade de notaas
	var (
		nota [numNotas]float64 // array de notas
		soma float64
	)
	//parte 1 - leitura
	for i := 0; i < numNotas; i++ {
		f.Printf("Informe a nota %d: ", i)
		f.Scan(&nota[i]) //armazena nota no array
		soma += nota[i]  //calcula a soma das notas
	}
	var media float64 = soma / 3

	f.Printf("media: %.2f\n\n", media)

	for i, v := range nota {
		f.Printf("Nota %d = %.2f\n", i, v)
		if nota[i] >= 6 {
			f.Printf("acima da media com nota: %.2f\n\n", v)
		} else {
			f.Printf("abaixo da media e reprovado com  %.2f\n\n", v)
		}
	}
	f.Printf("Soma da notas: %.2f ", soma)
}
