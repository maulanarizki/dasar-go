package main

import "fmt"

func operator() {
	// Tanda	Penjelasan
	// +	penjumlahan
	// -	pengurangan
	// *	perkalian
	// /	pembagian
	// %	modulus / sisa hasil pembagian

	var value = (((2 + 6) % 3) * 4 - 2) / 3

	fmt.Println(value)

	// Tanda	Penjelasan
	// ==	apakah nilai kiri sama dengan nilai kanan
	// !=	apakah nilai kiri tidak sama dengan nilai kanan
	// <	apakah nilai kiri lebih kecil daripada nilai kanan
	// <=	apakah nilai kiri lebih kecil atau sama dengan nilai kanan
	// >	apakah nilai kiri lebih besar dari nilai kanan
	// >=	apakah nilai kiri lebih besar atau sama dengan nilai kanan
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	// Tanda	Penjelasan
	// &&	kiri dan kanan
	// ||	kiri atau kanan
	// !	negasi / nilai kebalikan

	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}

func main() {
	operator()
}