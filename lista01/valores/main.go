package main

import "fmt"

func main() {
	var a, b, c, d, e, f, g, h, i, j int
	fmt.Scan(&a, &b, &c, &d, &e, &f, &g, &h, &i, &j)
	media := (a + b + c + d + e + f + g + h + i + j) / 10
	fmt.Println("A média dos números de 1 a 10 é:", media)
}
