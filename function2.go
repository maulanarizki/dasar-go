package main

import "fmt"
import "math"

// func calculate(d float64) (float64, float64) {
//     // hitung luas
//     var area = math.Pi * math.Pow(d / 2, 2)
//     // hitung keliling
//     var circumference = math.Pi * d

//     // kembalikan 2 nilai
//     return area, circumference
// }

func calculate(d float64) (area float64, circumference float64) {
    area = math.Pi * math.Pow(d / 2, 2)
    circumference = math.Pi * d

    return
}

func main() {
    var diameter float64 = 15
    var area, circumference = calculate(diameter)

    fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
    fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

// Fungsi math.Pow() digunakan untuk memangkat nilai. math.Pow(2, 3) berarti 2 pangkat 3, hasilnya 8. Fungsi ini berada dalam package math
// math.Pi adalah konstanta bawaan package math yang merepresentasikan Pi atau 22/7