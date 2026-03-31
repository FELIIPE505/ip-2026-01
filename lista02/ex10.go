package main

import (
	f "fmt"
)

func main() {
	var destino, idavolta int
	f.Printf("DIGITE UM 1 PARA REGIAO NORTE: \n2 PARA REGIAO NORDESTE: \n3 PARA REGIAO CENTRO-OESTE: \n 4PARA REGIAO SUL: \n 1 PARA IDA E VOLTA: \nE 2 PARA SÓ IDA: ")
	f.Scan(&destino, &idavolta)

	if destino == 1 {
		if idavolta == 1 {
			f.Printf("O VALOR DA PASSAGEM É 900")
		} else {
			f.Printf("O VALOR DA PASSAGEM É 500")
		}
	} else if destino == 2 {
		if idavolta == 1 {
			f.Printf("O VALOR DA PASSAGEM É 600.50")
		} else {
			f.Printf("O VALOR DA PASSAGEM É 350")
		}
	} else if destino == 3 {
		if idavolta == 1 {
			f.Printf("O VALOR DA PASSAGEM É 600")
		} else {
			f.Printf("O VALOR DA PASSAGEM É 350")
		}
	} else if destino == 2 {
		if idavolta == 4 {
			f.Printf("O VALOR DA PASSAGEM É 550")
		} else {
			f.Printf("O VALOR DA PASSAGEM É 300")
		}
	}
}
