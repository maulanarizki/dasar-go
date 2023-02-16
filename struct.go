package main

import "fmt"

type student struct {
    name string
    grade int
}

type person1 struct {
    name string
    age  int
}

type student1 struct {
    grade int
    person1
}

type person2 struct {
    name string
    age  int
}

type student2 struct {
    person2
    age   int
    grade int
}

func struct1(){
	var s1 student
    s1.name = "john wick"
    s1.grade = 2

    fmt.Println("name  :", s1.name)
    fmt.Println("grade :", s1.grade)

	var s2 = student{}
	s2.name = "wick"
	s2.grade = 2

	var s3 = student{"ethan", 2}

	var s4 = student{name: "jason"}

	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)
	fmt.Println("student 4 :", s4.name)

	var s5 = student{name: "wayne", grade: 2}
	var s6 = student{grade: 2, name: "bruce"}
	
	fmt.Println("student 5 :", s5.name)
	fmt.Println("student 6 :", s6.name)
}

func struct2()  {
	var s1 = student{name: "wick", grade: 2}

	var s2 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)

	s2.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)
}

func struct3()  {
	var s1 = student1{}
    s1.name = "wick"
    s1.age = 21
    s1.grade = 2

    fmt.Println("name  :", s1.name)
    fmt.Println("age   :", s1.age)
    fmt.Println("age   :", s1.person1.age)
    fmt.Println("grade :", s1.grade)
}

func struct4()  {
	var s1 = student2{}
    s1.name = "wick"
    s1.age = 21        // age of student
    s1.person2.age = 22 // age of person

    fmt.Println(s1.name)
    fmt.Println(s1.age)
    fmt.Println(s1.person2.age)
}

func struct5()  {
	var p1 = person1{name: "wick", age: 21}
	var s1 = student1{person1: p1, grade: 2}

	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("grade :", s1.grade)
}

func struct6()  {
	var s1 = struct {
        person1
        grade int
    }{}
    s1.person1 = person1{"wick", 21}
    s1.grade = 2

    fmt.Println("name  :", s1.person1.name)
    fmt.Println("age   :", s1.person1.age)
    fmt.Println("grade :", s1.grade)
}

func slicestruct()  {
	type person struct {
		name string
		age  int
	}
	
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}
	
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}

func sliceanonymousstruct() {
	var allStudents = []struct {
		person1
		grade int
	}{
		{person1: person1{"wick", 21}, grade: 2},
		{person1: person1{"ethan", 22}, grade: 3},
		{person1: person1{"bond", 21}, grade: 3},
	}
	
	for _, student := range allStudents {
		fmt.Println(student.person1.name, "age is", student.person1.age, "grade is", student.grade)
	}
}

func declareanonymousvar()  {
	var student struct {
		person1
		grade int
	}
	
	student.person1 = person1{"wick", 21}
	student.grade = 2
	
	fmt.Println(student.person1.name, "age is", student.person1.age, "grade is", student.grade)
}

func nestedstruct()  {
	var studentall struct {
		maul struct {
			name string
			age  int
		}
		grade   int
		hobbies []int
	}
	
	studentall.maul = person1{"wick", 21}
	studentall.grade = 2
	studentall.hobbies = []int{1,4,3,4,2}

	// alias
	type Person struct {
		name string
		age  int
	}
	type People = Person

	var p1 = Person{"wick", 21}
	fmt.Println(p1)

	var p2 = People{"wick", 21}
	fmt.Println(p2)
	fmt.Println(studentall)
}

func main()  {
	struct1()
	struct2()
	struct3()
	struct4()
	struct5()
	struct6()
	slicestruct()
	sliceanonymousstruct()
	declareanonymousvar()
	nestedstruct()
}