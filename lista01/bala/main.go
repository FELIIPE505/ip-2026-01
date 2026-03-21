package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	slice := []int{a, b, c}
	sort.Ints(slice)

	balamenor := slice[0] * 3
	balainter := slice[1] * 2
	balamaior := slice[2] * 1
	total := balamenor + balainter + balamaior
	fmt.Printf("Total de balinhas a comprar: %d\n", total)
	fmt.Printf("%d balinhas para o neto de idad %d\n", balamenor, slice[0])
	fmt.Printf("%d balinhas para o neto de idade %d\n", balainter, slice[1])
	fmt.Printf("%d balinhas para o neto de idade %d\n", balamaior, slice[2])
}
