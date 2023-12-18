package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	a.SetString("4294967296", 10) // 2^32
	b.SetString("268435456", 10)  // 2^28

	// Умножение
	product := new(big.Int)
	product.Mul(a, b)
	fmt.Printf("Умножение: %s * %s = %s\n", a.String(), b.String(), product.String())

	// Деление (предполагается, что b не равно 0)
	division := new(big.Int)
	division.Div(a, b)
	fmt.Printf("Деление: %s / %s = %s\n", a.String(), b.String(), division.String())

	// Сложение
	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Printf("Сложение: %s + %s = %s\n", a.String(), b.String(), sum.String())

	// Вычитание
	difference := new(big.Int)
	difference.Sub(a, b)
	fmt.Printf("Вычитание: %s - %s = %s\n", a.String(), b.String(), difference.String())
}
