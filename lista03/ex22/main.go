package main

import f "fmt"

func main() {
	var s float64
	for i := 1; i <= 37; i++ {
		s = s + (38-float64(i))*(39-float64(i))/float64(i)
	}
	f.Println("S =", s)
}
