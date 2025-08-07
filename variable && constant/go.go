package main

import "fmt"

// var name = "smithcheal"

func main() {
	// name_3 := "leowave"
	// // check := "gs"
	// fmt.Println(reflect.TypeOf(name_3))

	// var a, b, c string
	// a = "chuks"
	// b = "smithcheal"
	// c = "chidera"
	// fmt.Println(a, b, c)

	//multi variable dec
	// var (
	// 	chidera = "chidera"
	// 	smith   = "smithcheal"
	// )

	// fmt.Println(chidera, smith)
	// var num1, num2, num3 int
	// num1 = 2
	// num2 = 4
	// num3 = 70
	//

	// constant
	// const name string = "kodex"
	// const name_5 string
	// name_5 = "chuks"
	// fmt.Println(name_5)

	// const (
	// 	num1 = 3
	// 	num2 = 5
	// 	num3 = 7
	// )

	// fmt.Println(num1, num2, num3)
	// fmt.Printf("%T", num1)

	// iota
	// var (
	// 	num1 = iota + 5
	// 	num2
	// 	num3
	// )
	// fmt.Println(num1)
	// fmt.Println(num2)
	// fmt.Println(num3)
	const (
		_ = iota + 1
		num2
		num3
	)

	fmt.Println(num2)
	fmt.Println(num3)
}

//variable and constanst
// var
// bool
// :=
// iota
// In Go, **iota** is not an acronym — it doesn't have a "full meaning" like each letter standing for something.
// Instead, iota is a predeclared identifier in Go used to create constant sequences — mostly numeric — in a clean and automatic way.
// const (
//   _ = iota      // skip 0
//   Bronze        // 1
//   Silver        // 2
//   Gold          // 3
// )
// typeof fmt.Printf("%T", ...)
// reflect
// checking type of text
