package main

import f "fmt"

func main() {
	var n int
	f.Println("Qual termo deve ser calculado?")
	f.Scan(&n)

	termo := n * n

	f.Printf("O termo %d é ", termo)
}
