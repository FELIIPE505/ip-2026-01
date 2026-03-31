package main

import "fmt"

func main() {
	var n, m int
	fmt.Printf("Digite um numero: ")
	fmt.Scan(&n)
	fmt.Printf("Digite outro numero: ")
	fmt.Scan(&m)

	z := n + m

	if z > 20 {
		fmt.Printf("%d\n", z+8)
	} else {
		fmt.Printf("%d\n", z-5)
	}
}
