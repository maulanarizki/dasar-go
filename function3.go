// Go mengadopsi konsep variadic function atau pembuatan fungsi dengan parameter sejenis yang tak terbatas. Maksud tak terbatas di sini adalah jumlah parameter yang disisipkan ketika pemanggilan fungsi bisa berapa saja.
package main

import "fmt"
import "strings"

func main() {
    // var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg = calculate(numbers...)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)

    var hobbies = []string{"sleeping", "eating"}
    yourHobbies("wick", hobbies...)

	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
    var total int = 0
    for _, number := range numbers {
        total += number
    }

    var avg = float64(total) / float64(len(numbers))
    return avg
}

func yourHobbies(name string, hobbies ...string) {
    var hobbiesAsString = strings.Join(hobbies, ", ")

    fmt.Printf("Hello, my name is: %s\n", name)
    fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}