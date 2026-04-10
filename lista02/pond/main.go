package main

import (
	f "fmt"
)

func main() {
	var p1 float64
	var p2 float64
	f.Print("digite a pontuação 1: ")
	f.Scan(&p1)
	f.Print("digite a pontuação 2: ")
	f.Scan(&p2)

	mediap := ((p1 * 2) + (p2 * 3)) / 5
	if p1 >= 0.0 && p1 <= 10.0 && p2 >= 0.0 && p2 <= 10.0 {
		if mediap >= 7.0 {
			f.Print("Vitorioso")
		} else if mediap <= 3.0 {
			f.Print("Derrotado")
		} else {
			f.Print("Julgamento Final")
		}
	} else {
		f.Print("FORA DAS RESTRIÇÕES")
	}
}
