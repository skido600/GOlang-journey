package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func main() {
	fmt.Println("math and number")
	// maths and numbers
	// var number int = 44
	// num := 55
	// var feet float32 = 0.5

	// fmt.Println(feet)
	// fmt.Println(reflect.TypeOf(num))
	// fmt.Println(number)

	//math package
	var x int = 2
	var y int = 5
	fmt.Println(x - y)
	fmt.Println(x + y)

	//modules %%
	fmt.Println(math.Mod(18, 5))
	fmt.Println(18 % 5)
	//sortig
	name := []int{6, 4, 3, 1, 2}
	sort.Ints(name)

	fmt.Println(name)
	//check squareroot
	var num1 float64 = 25
	//check squareroot
	// var num2 int = 25
	fmt.Println(math.Sqrt(num1))

	//pow
	// var num2 = 25
	// fmt.Println(math.Pow(5, 2))

	//absolute
	fmt.Println(math.Abs(-5))
	fmt.Println(math.Floor(904.03))

	// rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100) + 1)
	//complex numbers
	c := complex(1, 3)
	fmt.Println(c)
	fmt.Println(real(c))
	fmt.Println(imag(c))
}
