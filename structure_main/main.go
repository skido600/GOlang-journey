package main

import "fmt"

//private

// modification on struct using

type scores struct {
	english int
	math    int
	local   int
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

	//make allow u to create empty map and assign it later
	//eg let see the differencis
	//map
	var names = map[string]string{
		"student1": "chuks",
		"student2": "leowave",
	}
	//accessing in map
	fmt.Println(names["student1"])
	//deleting using in map
	delete(names, "student1")
	//adding
	names["Ada"] = "cartel"
	//looping over map you cannot use for loop len in map map only use range to loop for
	for k, v := range names {
		fmt.Println("key", k, "-", "value", v)
	}
	//make make and map are the same thing in map u create and assign direct but in  make u just create empty map and assign later

}
