package main

import (
	"fmt"
)

func main() {
	// 	var a = 15 + 25
	// 	fmt.Println(a)

	// 	var (
	// 		sum1 = 100 + 50
	// 		sum2 = sum1 + 250
	// 		sum3 = sum2 + sum2
	// 	)
	// 	fmt.Println(sum3)

	// 	var x = 10
	// 	x += 5
	// 	fmt.Println(x)

	// logical operation
	// && and
	// || or
	// ! not
	// if statement
	// if condition {

	// // }
	// if 20 > 18 {
	// 	fmt.Println("20 is greater than 18")
	// }

	// 	switch expression {
	// case x:
	//    // code block
	// case y:
	//    // code block
	// case z:
	// ...
	// default:
	//    // code block
	// }

	// day := 5

	// switch day {
	// case 1, 3, 5:
	// 	fmt.Println("Odd weekday")
	// case 2, 4:
	// 	fmt.Println("Even weekday")
	// case 6, 7:
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Invalid day of day number")
	// }

	// for i := 0; i < 5; i++ {
	// 	// fmt.Println(i)
	// 	if i == 2 {
	// 		fmt.Println("Skipping", i)
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// num1 := 1

	// if condition {

	// }
	//if else
	// if num1 < 2 {
	// 	fmt.Println("greater than")
	// } else {
	// 	fmt.Println("less than")
	// }
	//if else statement
	// num3 := 100
	// if num3 > 2 {
	// 	fmt.Println("greater than")
	// } else if num3 >= 100 {
	// 	fmt.Println("num num")
	// } else {
	// 	fmt.Println("default")
	// }

	//logical statement
	// && and
	// || or
	// ! not

	//and &&
	// name := false
	// name2 := true
	// if name && name2 {
	// 	fmt.Println("we are having class today")
	// } else {
	// 	fmt.Println("no class")
	// }

	//|| or
	// name3 := false
	// name4 := true
	// if name3 || name4 {
	// 	fmt.Println("we are having class today")
	// } else {
	// 	fmt.Println("no class")
	// }

	// ! not
	// var iphone bool = false

	// if !iphone {
	// 	fmt.Println("leo and smith are not using an iphone")
	// } else {
	// 	fmt.Println("we are  using iphone")
	// }

	// 	switch expression {
	// case x:
	//    // code block
	// case y:
	//    // code block
	// case z:
	// ...
	// default:
	//    // code block
	// }
	// day := "wednesday"

	// switch day {
	// case "tuesday":
	// 	fmt.Println("we are on tuesday")
	// case "friday":
	// 	fmt.Println("we are on friday")
	// case "wednesday":
	// 	fmt.Println("we are on wednesday")
	// default:
	// 	fmt.Println("we are not on anyday")
	// }

	//multiple switch
	// week := "week2"

	// switch week {
	// case "week1", "week2":
	// 	fmt.Println("we are on week")
	// default:
	// 	fmt.Println("we are not on anyweek")
	// }

	//loop

	// for i := 0; i <= 3; i++ {
	// 	fmt.Println("2x", i, "=", i*2)
	// }
	for i := 1; i <= 5; i++ {

		if i == 1 {
			continue
		}
		fmt.Println(i)
	}

}
