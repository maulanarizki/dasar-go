// package main

// import "fmt"
// import "strings"

// func main()  {
// 	var names = []string{"Maulana", "Rizki"}
// 	printMessage("Halo", names)
// }

// func printMessage(message string, arr []string) {
//     var nameString = strings.Join(arr, " ")
//     fmt.Println(message, nameString)
// }

// package main

// import (
//     "fmt"
//     "math/rand"
//     "time"
// )

// func main() {
//     rand.Seed(time.Now().Unix())
//     var randomValue int

//     randomValue = randomWithRange(2, 10)
//     fmt.Println("random number:", randomValue)
//     randomValue = randomWithRange(2, 10)
//     fmt.Println("random number:", randomValue)
//     randomValue = randomWithRange(2, 10)
//     fmt.Println("random number:", randomValue)
// }

// func randomWithRange(min, max int) int {
//     var value = rand.Int() % (max - min + 1) + min
//     return value
// }

// Khusus untuk fungsi yang tipe data parameternya sama, bisa ditulis dengan gaya yang unik. Tipe datanya dituliskan cukup sekali saja di akhir. Contohnya bisa dilihat pada kode berikut.
// func nameOfFunc(paramA type, paramB type, paramC type) returnType
// func nameOfFunc(paramA, paramB, paramC type) returnType

// func randomWithRange(min int, max int) int
// func randomWithRange(min, max int) int


package main

import "fmt"

func main() {
    divideNumber(10, 2)
    divideNumber(4, 0)
    divideNumber(8, -4)
}

func divideNumber(m, n int) {
    if n == 0 {
        fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
        return
    }

    var res = m / n
    fmt.Printf("%d / %d = %d\n", m, n, res)
}