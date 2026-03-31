package main

import (
	f "fmt"
	"math"
)

func main() {
	var n float64
	f.Print("INSIRA UM NÚMERO: ")
	f.Scan(&n)

	if n > 0 || n == 0 {
		raiz := math.Sqrt(n)
		f.Printf("A RAIZ QUADRADA DE %.2f É: %2.f\n", n, raiz)
	} else if n < 0 {
		qdrd := n * n
		f.Printf("O QUADRADO DE %.2f É: %.2f", n, qdrd)
	} else {
		f.Printf("INPUT INVÁLIDO")
	}
}
