package main

import f "fmt"

func impar(n int) int {
	var x int
	x = 1 + (n-1)*2

	return x
}
func main() {
	var h float64
	h = 0

	for i := 1; i <= 50; i++ {
		h = h + float64(impar(i))/float64(i)
	}
	f.Println("H = ", h)

}
