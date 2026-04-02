package main

import (
	f "fmt"
)

func main() {
	var x int
	f.Scan(&x)

	fato := fatorial(x)
	f.Printf("%.d", fato)

}

func fatorial(x1 int) int {
	resultado := 1
	for i := 2; i <= x1; i++ {
		resultado *= i
	}
	return resultado
}
