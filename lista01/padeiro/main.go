package main

import "fmt"

func main() {
	var qtdkgfarinha float64
	var valorkgfarinha float64
	var qtdkgfermento float64
	var valorkgfermento float64
	var qtdkgwatts float64
	var valorkgwatts float64
	var percentualpelopao float64
	fmt.Scan(&qtdkgfarinha, &valorkgfarinha, &qtdkgfermento, &valorkgfermento, &qtdkgwatts, &valorkgwatts, &percentualpelopao)

	finalfarinha := qtdkgfarinha * valorkgfarinha
	finalfermento := qtdkgfermento * valorkgfermento
	finalwatts := qtdkgwatts * valorkgwatts
	finalcusto := (finalfarinha + finalfermento + finalwatts) * (1 + percentualpelopao/100)
	lucro := finalcusto * 0.3
	valorfinal := finalcusto + lucro
	fmt.Printf("PRECO DE CUSTO = %.2f\n", finalcusto)
	fmt.Printf("PRECO DE VENDA = %.2f\n", valorfinal)
}
