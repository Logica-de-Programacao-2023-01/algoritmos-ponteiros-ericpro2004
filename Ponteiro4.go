package main

import "fmt"

func updateLastTwoDigits(ptr *int) {
	num := *ptr
	digit1 := num % 10
	num /= 10
	digit2 := num % 10
	*ptr = digit1 + digit2
}

func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)

	updateLastTwoDigits(&num)

	fmt.Println("O novo valor é:", num)
}
