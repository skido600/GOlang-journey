package main

//string and method
import (
	"fmt"
	"strings"
	s "strings"
)

func main() {
	// variable 'name' contains the string "wave"

	name := "WAVE"
	var greeting string = "hello world"
	fmt.Println(greeting)
	greeting = "cartel"

	// Contains:e
	fmt.Println(s.Contains(name, "w"))

	// Count: count how many times
	fmt.Println(s.Count(name, "w"))

	// HasPrefix: check if 'name' starts
	fmt.Println(s.HasPrefix(name, "w"))

	// HasSuffix: check if 'name' ends
	fmt.Println(s.HasSuffix(name, "w"))
	//check the number of index
	fmt.Println(s.Index(name, "e"))
	//convert to lowercase
	fmt.Println(s.ToLower(name))
	//to uppercase
	alphabet := "smithcheal"
	fmt.Println(s.ToUpper(alphabet))

	strin()

}

func strin() {
	// fmt.Println("string \n \tday")
	//day4
	//string immutability
	var name string = "leowave"
	fmt.Println(name)
	name = "leowave_2"
	//concatination quote concatination
	name_one := "is learning golang string"
	name_two := "smith"
	fmt.Println(name_two + " " + name_one)
	// format specifiers
	// %s → string
	// %d → integer
	// %f → float
	// num1 := 100
	// name_three := "kelvin"
	// var float_number float64 = 0.4

	// fmt.Printf("num1:%d,name_three:%s float %f", num1, name_three, float_number)

	//slicing indexing
	var datavar string = "we are learning"
	newstring := "j" + datavar[1:]
	fmt.Println(newstring)
	//slicing
	// sliceme[start:end]
	var sliceme string = "smith"
	fmt.Println(sliceme[0:1])

	//escape character
	// escape := "wave\nmm"
	// escape_2 := "smith\tcheal"
	// fmt.Println(escape_2)

	//string method
	//contains
	var text string = " hello world, welcome to the world"
	fmt.Println(strings.Contains(text, "w"))
	//touppercase
	fmt.Println(strings.ToUpper(text))
	//tolowwercase
	fmt.Println(strings.ToLower(text))
	//count
	fmt.Println(strings.Count(text, "a"))
	//to check if text startwith a in js we start with
	fmt.Println(strings.HasPrefix(text, "w"))
	//to check if text endwith a in js we start with
	fmt.Println(strings.HasSuffix(text, "g"))
	//split convert
	fmt.Println(strings.Split(text, ","))
	//replace
	fmt.Println(strings.Replace(text, "world", "war", 2))

}
