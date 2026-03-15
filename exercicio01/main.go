package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	mg := (a + b + c) / 3

	fmt.Printf("MEDIA = %.2f\n", mg)
	if mg >= 6 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
