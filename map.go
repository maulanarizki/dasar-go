package main

import "fmt"

func map1() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei",     chicken["mei"])     // mei 0

	// var data map[string]int
	var data = map[string]int{}
	data["one"] = 1
	// akan muncul error!

	data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	var chicken3 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	for key, val := range chicken1 {
		fmt.Println(key, "  \t:", val)
	}
	for key, val := range chicken2 {
		fmt.Println(key, "  \t:", val)
	}
	for key, val := range chicken3 {
		fmt.Println(key, "  \t:", val)
	}

	fmt.Println(len(chicken1)) // 2
	fmt.Println(chicken1)

	delete(chicken1, "januari")

	fmt.Println(len(chicken1)) // 1
	fmt.Println(chicken1)

	var chicken4 = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken4["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// var chickens = []map[string]string{
	// 	map[string]string{"name": "chicken blue",   "gender": "male"},
	// 	map[string]string{"name": "chicken red",    "gender": "male"},
	// 	map[string]string{"name": "chicken yellow", "gender": "female"},
	// }

	var chickens = []map[string]string{
		{"name": "chicken blue",   "gender": "male"},
		{"name": "chicken red",    "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	var data1 = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	
	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	// fmt.Println(data1)

	for _, data1 := range data1 {
		fmt.Println(data1)
	}

}

func main()  {
	map1()
}