package main

import "fmt"

func checkParImpar(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0 // atualiza para 0 se for par
	} else {
		*ptr = 1 // atualiza para 1 se for ímpar
	}
}

func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)

	checkParImpar(&num)

	fmt.Println("O valor atualizado é:", num)
}
