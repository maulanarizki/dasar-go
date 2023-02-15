package main

import "fmt"

func slice()  {
	// Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama
	var fruits = []string{"apple", "grape", "banana", "melon"} // slice ngak ada ditentukan panjang isinya 
	fmt.Println(fruits[0]) // "apple"

	var fruitsA = []string{"apple", "grape"}      // slice
	var fruitsB = [2]string{"banana", "melon"}    // array
	var fruitsC = [...]string{"papaya", "grape"}  // array
	fmt.Println(fruitsA) // "apple"
	fmt.Println(fruitsB) // "apple"
	fmt.Println(fruitsC) // "apple"

	var newFruits = fruits[0:2]

	fmt.Println(newFruits) // ["apple", "grape"]

	// Kode	Output	Penjelasan
	// fruits[0:2]	[apple, grape]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
	// fruits[0:4]	[apple, grape, banana, melon]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
	// fruits[0:0]	[]	menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
	// fruits[4:4]	[]	menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
	// fruits[4:0]	[]	error, pada penulisan fruits[a:b] nilai a harus lebih kecil atau sama dengan b
	// fruits[:]	[apple, grape, banana, melon]	semua elemen
	// fruits[2:]	[banana, melon]	semua elemen mulai indeks ke-2
	// fruits[:2]	[apple, grape]	semua elemen hingga sebelum indeks ke-2

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	// funsi len 
	fmt.Println(len(fruits)) // 4

	// Fungsi cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice. Nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan len, tapi bisa berubah seiring operasi slice yang dilakukan. Agar lebih jelas, silakan disimak kode berikut.

	fmt.Println(len(fruits))  // len: 4
	fmt.Println(cap(fruits))  // cap: 4

	var a1Fruits = fruits[0:3]
	fmt.Println(len(a1Fruits)) // len: 3
	fmt.Println(cap(a1Fruits)) // cap: 4

	var b1Fruits = fruits[1:4]
	fmt.Println(len(b1Fruits)) // len: 3
	fmt.Println(cap(b1Fruits)) // cap: 3

	// append array/slice 
	var cFruits = append(fruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	var b2Fruits = fruits[0:2]

	fmt.Println(cap(b2Fruits)) // 3
	fmt.Println(len(b2Fruits)) // 2

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(b2Fruits) // ["apple", "grape"]

	var c2Fruits = append(b2Fruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	fmt.Println(b2Fruits) // ["apple", "grape"]
	fmt.Println(c2Fruits) // ["apple", "grape", "papaya"]

	// copy() = Fungsi copy() digunakan untuk men-copy elements slice pada src (parameter ke-2), ke dst (parameter pertama) 
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	// 3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya. Cara menggunakannnya yaitu dengan menyisipkan angka kapasitas di belakang, seperti fruits[0:1:1]. Angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice yang akan di slicing.
	var a3Fruits = fruits[0:2]
	var b3Fruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(a3Fruits)      // ["apple", "grape"]
	fmt.Println(len(a3Fruits)) // len: 2
	fmt.Println(cap(a3Fruits)) // cap: 3

	fmt.Println(b3Fruits)      // ["apple", "grape"]
	fmt.Println(len(b3Fruits)) // len: 2
	fmt.Println(cap(b3Fruits)) // cap: 2
}

func main()  {
	slice()
}