package main

import (
	"fmt"
	"math"
)

func updateAverageWithPi(ptr *float64) {
	pi := math.Pi
	currentValue := *ptr
	average := (currentValue + pi) / 2
	*ptr = average
}

func main() {
	var num float64
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	updateAverageWithPi(&num)

	fmt.Println("O novo valor é:", num)
}
