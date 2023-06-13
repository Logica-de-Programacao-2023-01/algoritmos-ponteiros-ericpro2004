package main

import (
	"fmt"
)

func modifyValue(ptr *int) {
	*ptr = 42
}

func main() {
	num := 10
	ptr := &num

	fmt.Println("Antes da modificação:", num)

	modifyValue(ptr)

	fmt.Println("Após a modificação:", num)

	ptr = nil
}
