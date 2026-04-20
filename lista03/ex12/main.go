package main

import (
	f "fmt"
	m "math"
)

func fatorial(n int) int {
	var P int = 1

	for i := 1; i <= n; i++ {
		P = P * i
	}
	return P
}
func main() {
	var x, s float64

	f.Println("Insira um valor X: ")
	f.Scan(&x)
	s = x
	for i := 1; i < 20; i++ {
		s = s + x*m.Pow(-1, float64(i))/float64(fatorial(i))
	}
	f.Println("S = ", s)
}
