package main

import "fmt"

func updateValue(ptr *int, n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	*ptr = sum
}

func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)

	var n int
	fmt.Print("Digite o valor de n: ")
	fmt.Scan(&n)

	updateValue(&num, n)

	fmt.Println("O valor atualizado é:", num)
}
