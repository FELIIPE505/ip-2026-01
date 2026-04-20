package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	crescente := []int{a, b, c}
	sort.Ints(crescente)
	fmt.Println(crescente[2])
}
