package main

import (
	f "fmt"
)

func main() {
	l := []int{1, 2, 3, 4, 5}
	x := 3
	pos := buscaseque(l, x)
	if pos != -1 {
		f.Printf("Elemento %d encontrado na posição %d\n", x, pos)
	}
}

func buscaseque(l []int, x int) int {
	if n := len(l); n > 0 {
		for i := 0; i < n; i++ {
			if l[i] == x {
				return i
			}
		}
		f.Print("Elemento não encontrado")
	}
	return -1
}

/* l := []int{1, 2, 3}
x := 3

pos :=  busca(l, x)
if pos != -1 {
	f.Print(""Elemento %d encontrado na posição %d\n", x, pos")
}

func busca (l []int, x int) int{
	if n := len(1); n > 0 {
		for i := 0; i < n; i++{
			if l[i] == x {
				return i
			}
		}
	}
	return -1
} */
