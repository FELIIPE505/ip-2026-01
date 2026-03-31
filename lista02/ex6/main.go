package main

import (
	f "fmt"
)

func main() {
	var n, m int
	f.Println("DIGITE UM NUMERADOR E UM DENOMINADOR : ")
	f.Scan(&n, &m)

	if n%m == 0 && n > 1 || n < 0 {
		f.Printf("O NÚMERO %d É DIVISIVEL POR %d", n, m)
	} else if n == 0 && m == 0 || n == 0 || m == 0 {
		f.Printf("NÚMERO INVÁLIDO")
	} else {
		f.Printf("O NÚMERO NÃO É DIVISIVEL POR %d", m)
	}
}
