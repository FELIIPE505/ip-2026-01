package main

import (
	f "fmt"
	m "math"
)

func main() {

	var a float64

	if a < 0.0 || a > 6.3 {
		f.Println("Valor inserido inválido")
	}

	for k := a; k <= 6.3; k += 0.1 {
		y := k - m.Pow(k, 3)/6 + m.Pow(k, 5)/120 - m.Pow(k, 7)/5040
		if y <= 1 && y >= -1 {
			f.Printf("Sen(%f) = %f", k, y)
			f.Println()
		} else {
			break
		}
	}
}
