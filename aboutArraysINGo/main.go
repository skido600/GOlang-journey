package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	//array
	var age [4]int = [4]int{1, 2, 3, 4}
	age2 := [4]int{1, 2, 3, 4}

	fmt.Println(age)
	fmt.Println(age2, len(age2))
	//array string
	names := [3]string{"dog", "cat", "monkey"}

	//lenght len()
	//updating array
	// names[0] = "bulldog"
	// fmt.Println(names)

	//slice
	scores := []int{90, 40, 30}
	scores[0] = 10

	//to add at back
	scores = append(scores, 0)
	fmt.Println(scores)
	//ranges
	range1 := names[1:3]
	range2 := names[:2]
	range3 := names[:3]

	fmt.Println(range2)

	fmt.Println(range3)
	//range appeding
	range1 = append(range1, "key")
	fmt.Println(range1)
	//mapping
	menu := map[string]float64{
		"yam":   12.99,
		"beans": 20.8,
	}
	fmt.Println(menu["yam"], len(menu))
	//looping
	for key, v := range menu {

		fmt.Println(key, v)

	}
	//int part
	menu2 := map[int]string{
		90: "akpu",
		40: "beans",
	}
	menu2[40] = "rice"
	for key, value := range menu2 {
		fmt.Println(key, "-", value)

	}
	//deleting
	delete(menu2, 90)
	fmt.Println(menu2)

	nums := make([]int, 3, 5)
	fmt.Println(nums)

}
