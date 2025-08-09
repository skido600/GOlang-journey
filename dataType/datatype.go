package main

import (
	"fmt"
	"reflect"
)

func main() {
	// talking about data type in golang
	//string
	//int
	//boolean
	///float

	//string
	var name_1 string = "mr leo"
	var name_2 = "smithcheal"
	var name_3 string
	fmt.Println(reflect.TypeOf(name_3))
	fmt.Println(name_1, name_2)
	mr_gabriel := "mr gab"
	fmt.Println(mr_gabriel)
	//int
	// var num_1 int = 44
	// fmt.Printf("%T", num_1)
	// var intvarone int64
	// fmt.Println(intvarone)

	//bool

	var isAdmin bool = true
	fmt.Println(isAdmin)

	var floatinvar float64 = 5124623583973.3
	fmt.Println(floatinvar)
}
