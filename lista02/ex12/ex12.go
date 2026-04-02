package main

import (
	f "fmt"
)

func main() {
	var idade int
	f.Println("Digite sua idade: ")
	f.Scan(&idade)

	if idade >= 0 && idade <= 2 {
		f.Printf("recém nascido")
	} else if idade >= 3 && idade <= 11 {
		f.Printf("criança")
	} else if idade >= 12 && idade <= 19 {
		f.Printf("adolescente")
	} else if idade >= 20 && idade <= 55 {
		f.Printf("adulto")
	} else if idade > 55 {
		f.Printf("idoso")
	}
}
