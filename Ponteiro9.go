package main

import (
	"fmt"
)

type Livro struct {
	Título string
	Autor  string
	Preço  float64
}

func aplicarDesconto(livro *Livro) {
	desconto := 0.1 // 10% de desconto
	livro.Preço -= livro.Preço * desconto
}

func main() {
	livro := Livro{
		Título: "Aventuras no Espaço",
		Autor:  "João Silva",
		Preço:  50.0,
	}

	fmt.Printf("Preço antes do desconto: R$ %.2f\n", livro.Preço)

	aplicarDesconto(&livro)

	fmt.Printf("Preço após o desconto: R$ %.2f\n", livro.Preço)
}
