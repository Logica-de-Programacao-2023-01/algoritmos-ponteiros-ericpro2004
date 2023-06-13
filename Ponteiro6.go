package main

import (
	"fmt"
)

type Livro struct {
	titulo string
	autor  string
}

func updateTitleIfAuthorIsAnonymous(livro *Livro) {
	if livro.autor == "Anônimo" {
		livro.titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		titulo: "Livro Exemplo",
		autor:  "Anônimo",
	}

	fmt.Println("Antes da atualização:", livro)

	updateTitleIfAuthorIsAnonymous(&livro)

	fmt.Println("Após a atualização:", livro)
}
