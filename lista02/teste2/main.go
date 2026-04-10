package main

import (
	f "fmt"
)

func main() {
	var S int
	f.Print("digite a quantidade de soldados:")
	f.Scan(&S)

	var D int
	f.Print("digite a quantidade de droids:")
	f.Scan(&D)

	cabe := S + D
	if S >= 1 && S <= 1000 && D >= 1 && D <= 1000 {
		if cabe <= 1000 {
			f.Print("S")
		} else {
			f.Print("N")
		}
	} else {
		f.Print("NUMERO INVALIDO")
	}
}
