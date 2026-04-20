package main

import f "fmt"

func primo(n int) bool {

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var n1, n2 int
	encontrou := false
	f.Println("Digite o valor do intervalo (N1 a N2):")
	f.Scan(&n1, &n2)

	for i := n1; i <= n2; i++ {
		if primo(i) {
			f.Printf("%d ", i)
			encontrou = true
		}
	}
	if !encontrou {
		f.Println("Não há número primo no intervalo")
	}
}
