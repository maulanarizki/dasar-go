// Di Go, fungsi bisa dijadikan sebagai tipe data variabel. Dari situ sangat memungkinkan untuk menjadikannya sebagai parameter juga

package main

import "fmt"
import "strings"
// import "os"

// func filter(data []string, callback func(string) bool) []string {
//     var result []string
// 	// fmt.Println(callback)
// 	// os.Exit(1)
//     for _, each := range data {
//         if filtered := callback(each); filtered {
//             result = append(result, each)
//         }
//     }
// 	// result = data
//     return result
// }

func filter(data []string, FilterCallback func(string) bool) []string {
    var result []string
	// fmt.Println(callback)
	// os.Exit(1)
    for _, each := range data {
        if filtered := FilterCallback(each); filtered {
            result = append(result, each)
        }
    }
	// result = data
    return result
}

func main() {
    var data = []string{"wick", "jason", "ethan"}
    var dataContainsO = filter(data, func(each string) bool {
        return strings.Contains(each, "o")
    })
    var dataLenght5 = filter(data, func(each string) bool {
        return len(each) == 5
    })

    fmt.Println("data asli \t\t:", data)
    // data asli : [wick jason ethan]

    fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
    // filter ada huruf "o" : [jason]

    fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
    // filter jumlah huruf "5" : [jason ethan]
}