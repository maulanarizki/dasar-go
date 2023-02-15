package main

import "fmt"

func array() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	// fmt.Println(names[0], names[1], names[2], names[3])

	// Menentukan jumlah slot array dr awal
	var name [4]string
	name[0] = "trafalgar"
	name[1] = "d"
	name[2] = "water"
	name[3] = "law"
	// names[4] = "ez" // baris kode ini menghasilkan error

	var fruits [4]string

	// cara horizontal
	fruits  = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	fruits  = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	var fruits2 = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits2); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits2[i])
	}

	var fruits3 = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits3 {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	var fruits4 = [4]string{"apple dsfsdf", "grape", "banana", "melon"}

	for _, fruit := range fruits4 {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	var fruits5 = make([]string, 2)
	fruits5[0] = "apple sdfdsf"
	fruits5[1] = "manggo"

	fmt.Println(fruits5)
}

func main() {
	array()
}