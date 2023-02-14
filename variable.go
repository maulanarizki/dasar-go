package main

import "fmt"

func variable() {
	// var <nama-variabel> <tipe-data> = <nilai>
	var firstname string = "Maulana"

	var lastname string
	lastname = "Rizki"

	maul := "mulai"
	maul = "uji"
	maul = "selesai"

	var fourth, fifth, sixth string = "empat", "lima", "enam"
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan keranjang sampah.
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "john", "wick"

	fmt.Println("Hello,", firstname, lastname, maul)
	fmt.Println(fourth, fifth, sixth)
	fmt.Println(seventh, eight, ninth)
	fmt.Println(one, isFriday, twoPointTwo, say)
	fmt.Println(name)
}

func main() {
	variable()
}