package main

import (
	f "fmt"
)

func main() {
	var x int
	f.Scan(&x)
	soma := 8 / (2 - x)
	f.Printf("f(%.d) = 8 / (2 - %.d) = %.d", x, x, soma)
}
