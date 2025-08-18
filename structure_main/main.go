package main

import (
	"fmt"
	"reflect"
)

// public
type Users struct {
	Name      string
	Age       int
	IsStudent bool
}

//private
// type users struct {
// 	name      string
// 	age       int
// 	isStudent bool
// }

// modification on struct using
func (p *Users) updateName(newname string, newage int) {
	p.Name = newname
	p.Age = newage

}

type party struct {
	name string
	age  int
}

func main() {
	//core
	// A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
	// While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.
	// A struct can be useful for grouping data together to create records.

	// To declare a structure in Go, use the type and struct keywords:
	// 	type struct_name struct {
	//   member1 datatype;
	//   member2 datatype;
	//   member3 datatype;
	//   ...
	// }

	// p1 := Person{"Leo", 21, "Lagos"}
	// p2 := Person{Name: "John", Age: 30}
	// var testing = Person{City: "anam"}
	// fmt.Println(testing)
	// fmt.Println(p2)
	// fmt.Println(p1)

	// //way two annoynimous struct
	// p3 := struct {
	// 	title string
	// 	age   int
	// }{title: "leowave blog", age: 200}

	// p4 := &p3
	// fmt.Println(p4.age)
	// fmt.Println(p4.title)

	// p4 := map[int]string{
	// 	"apple":  5,
	// 	"banana": 10,
	// }
	// fmt.Println(p4["apple"])
	// fmt.Println(p4["banana"])

	students := Users{Name: "leo", Age: 23, IsStudent: true}
	// students.isStudent = false
	students.updateName("smith", 20)
	//pointer cone
	// v := reflect.ValueOf(students)
	studentskodex := &students
	v := reflect.ValueOf(studentskodex)
	fmt.Println(v)

	//looping

	// for k, v := range studentskodex.name {

	// }

	// fmt.Println(students.name, students.isStudent)

	var parties = map[string]party{
		"obi": {
			name: "wave",
			age:  25,
		},
		"akpu": {
			name: "akpu",
			age:  12,
		},
	}

	for v, k := range parties {
		fmt.Println(v, "-", k.age)

	}

	fmt.Println(parties["obi"].age)
}
