/* package main
import (
	"fmt"
)
type Pessoa struct {
	nome      string
	idade     int
	profissao string
	salario   int
}
func main() {
	var pes1 Pessoa
	// Especificacao da pes1
	pes1.nome = "Monica"
	pes1.idade = 35
	pes1.profissao = "Engenheiro de Software"
	pes1.salario = 7000
	// Imprimir as informacoes da pessoa1
	fmt.Println("nome: ", pes1.nome)
	fmt.Println("idade: ", pes1.idade)
	fmt.Println("profissao: ", pes1.profissao)
	fmt.Println("salario: ", pes1.salario)
}
*/

package main

import (
	f "fmt"
	"strings"
)

type pessoa struct {
	nome      string
	altura    float64
	pesoideal float64
}

func main() {
	var pessoas []pessoa
	for {
		var p pessoa
		f.Print("Digite nome e altura, digite 'FIM' para encerrar: ")
		f.Scan(&p.nome)

		if strings.ToUpper(p.nome) == "FIM" {
			break
		}

		f.Print("Digite a altura (ex: 1.75): ")
		f.Scanln(&p.altura)

		p.pesoideal = (72.7 * p.altura) - 58.0

		pessoas = append(pessoas, p)
		f.Println("pessoa cadastrada")
	}

	f.Println("\n--- Lista de Pessoas e Pesos Ideais ---")
	for _, p := range pessoas {
		f.Printf("Nome: %s, Altura: %.2f m, Peso Ideal: %.2f kg\n", p.nome, p.altura, p.pesoideal)
	}
}
