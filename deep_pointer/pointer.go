package main

import "fmt"

// func greet(name string) {
// 	fmt.Println("hello" + name)

// }

// // divider
// func divide(c, d float64) (float64, error) {
// 	if c == 0 || d == 0 {
// 		return 0, errors.New("cannot divide by 0")
// 	}
// 	return c / d, nil
// }

// int add
// func add(a, b int) int {
// 	return a + b
// }

// type Toy interface {
// 	MakeSound() string
// }
// type Dog struct {
// }

// func (d Dog) MakeSound() string {
// 	return "wof"
// }
// func PlayWithToy(c Toy) {
// 	fmt.Println(c.MakeSound())

// }
// type Math interface {
// 	Area() float64
// }

// type Formular struct {
// 	Pi, K float64
// }

// func (i Formular) Area() float64 {
// 	return i.K + i.Pi
// }

type Dog struct {
	head string
	leg  string
	body int
}

func (i *Dog) Updatename(name int) {
	i.body = name
}

func main() {

	r := Dog{head: "leo", leg: "KODEX", body: 20}
	r.Updatename(3)
	fmt.Println(r)
	// r := Formular{Pi: 10.9, K: 5.0}
	// fmt.Println("Area:", r.Area())

	// fmt.Println("pointer")

	// greet("leo")

	// result := add(1, 2)
	// fmt.Println(result)
	// // i, j := 10, 20
	// c := 1.7
	// d := 1.0

	// result, err := divide(c, d)
	// if err != nil {
	// 	fmt.Println("errors", err)
	// } else {
	// 	fmt.Println(result)
	// }

	// PlayWithToy(Dog{})

	// var num1 int
	// _, err := fmt.Scan(&num1)
	// if err != nil {
	// 	fmt.Println("you get an error")
	// } else {

	// b := 30
	// c := *&b
	// // d := &b
	// fmt.Print(c)
	// In Go, we have two operators for pointers:
	//   1) &  → "address-of" operator (gives memory address of a variable)
	//   2) *  → "dereference" operator (go to that address and get/change the value)

	// &i means "the memory address of i"
	// &j means "the memory address of j"
	// These addresses are numbers in memory, like "house numbers" where the data lives.
	// fmt.Println(&i, &j) // e.g. 0xc000014098 0xc0000140a0

	// p is a pointer variable that stores the memory address of i
	// p := &i

	// Printing p will show the memory address where i is stored
	// fmt.Println("address of i:", p)

	// *p means "go to the address stored in p and fetch the value"
	// fmt.Println("value of i through pointer:", *p) // prints 10

	// If we change the value at *p, the original variable i will also change,
	// because *p directly accesses the real memory where i is kept.
	// *p = 77
	// fmt.Println("new value of i:", i) // 77

	// Summary:
	//   & → take the address of a variable
	//   * → follow the address to read or write the actual value

	// example
	// var num1 int = 20

	// num2 := &num1

	// fmt.Println(*num2 / 2)
	//function

	//function to greet

	//function to addd
}
