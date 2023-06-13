package main

import (
	"fmt"
)

type Conta struct {
	saldo   float64
	titular string
}

func adicionarValor(conta *Conta, valor float64) {
	conta.saldo += valor
}

func main() {
	conta := Conta{
		saldo:   100.0,
		titular: "João",
	}

	fmt.Println("Saldo antes da adição:", conta.saldo)

	valor := 50.0
	adicionarValor(&conta, valor)

	fmt.Println("Saldo após a adição:", conta.saldo)
}
