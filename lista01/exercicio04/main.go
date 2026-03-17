package main

import "fmt"

func main() {
	var sm float64
	fmt.Scan(&sm)
	var kw float64
	fmt.Scan(&kw)

	cKw := (sm * 0.7) / 100
	vKw := kw * cKw
	desconto := vKw - (vKw * 0.1)
	fmt.Printf("Custo por kW: R$%.2f\n", cKw)
	fmt.Printf("Custo do consumo: R$%.2f\n", vKw)
	fmt.Printf("Custo com desconto: R$%.2f\n", desconto)
}
