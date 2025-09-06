package main

import (
	"errors"
	"fmt"
)

func testingError(num int) (string, error) {
	if num < 3 {
		return "", errors.New("number is negative")
	}
	return "number is posivite", nil
}
func main() {
	win, err := testingError(2)
	if err != nil {
		fmt.Println(`error`, err)
	}
	fmt.Println(win)

}
