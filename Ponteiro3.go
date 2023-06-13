package main

import (
	"fmt"
)

func reverseString(ptr *string) {
	original := *ptr
	runes := []rune(original)

	// Inverter os caracteres da string
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*ptr = string(runes)
}

func main() {
	var text string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&text)

	reverseString(&text)

	fmt.Println("A string invertida é:", text)
}
