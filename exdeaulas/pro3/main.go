package main

import "fmt"

func main() {
	var n, desconto float64
	fmt.Scan(&n, &desconto)
	valorDesconto := n - (n * (desconto / 100))
	fmt.Printf("%.2f\n", valorDesconto)
}
