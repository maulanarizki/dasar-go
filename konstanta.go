package main

import "fmt"

func konstanta() {
	const firstName string = "Maulana"
	fmt.Print("halo ", firstName, "!\n")

	const lastName = "Rizki"
	fmt.Print("nice to meet you ", lastName, "!\n")

	// Perbedaan print line dan print
	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")

	const (
		square			= "kotak"
		b
		isToday bool	= true
		numeric uint8	= 1
		floatNum		= 2.2
	)

	fmt.Println(square)
	fmt.Println(b)
	fmt.Println(isToday)
	fmt.Println(numeric)
	fmt.Println(floatNum)

	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
	fmt.Println(satu, dua)
	fmt.Println(three, four)
}

func main() {
	konstanta()
}