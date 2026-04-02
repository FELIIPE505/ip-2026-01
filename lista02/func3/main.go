package main

import (
	f "fmt"
)

func main() {
	var x, y, z float64
	f.Scan(&x, &y, &z)

	m := media(x, y, z)
	f.Printf("média é = %.2f", m)
}

func media(x1, y1, z1 float64) float64 {
	return (x1 + y1 + z1) / 3
}
