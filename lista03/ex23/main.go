package main

import (
	f "fmt"
	m "math"
)

func main() {
	var n int
	var s float64
	f.Println("Insira a quantidade de termos que deve ser calculado")
	f.Scan(&n)
	s = 0
	for i := 0; i < n; i++ {
		s = m.Pow(-1, float64(i))*(1000-3*float64(i))/(float64(i)+1) + s
	}
	f.Println("O valor da expressão caçculada até o termo n é:", s)
}
