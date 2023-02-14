package main

import "fmt"

func tipedata() {
	// Tipe data	Cakupan bilangan
	// uint8	0 ↔ 255
	// uint16	0 ↔ 65535
	// uint32	0 ↔ 4294967295
	// uint64	0 ↔ 18446744073709551615
	// uint	sama dengan uint32 atau uint64 (tergantung nilai)
	// byte	sama dengan uint8
	// int8	-128 ↔ 127
	// int16	-32768 ↔ 32767
	// int32	-2147483648 ↔ 2147483647
	// int64	-9223372036854775808 ↔ 9223372036854775807
	// int	sama dengan int32 atau int64 (tergantung nilai)
	// rune	sama dengan int32

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var maul = `Nama saya 
	"Maulana Rizki".Salam kenal.Mari belajar "Golang".`

	fmt.Println(maul)
}

func main() {
	tipedata()
}