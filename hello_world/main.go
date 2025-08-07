// package declaration

// this my day1
package main

// import statement to import the format package necessary for IO operation like scanf ,printf,println
import (
	"fmt"
	"reflect"
)
const PI = 3.14;

// Single-line comment — just like JS

/*
   Multi-line comment — also like JS
*/

func main(){

// fmt.Print("hello")
//string
var myname string="chuks"
name2 :="leowave"
fmt.Println(myname)
fmt.Println(PI)
fmt.Println(reflect.TypeOf( myname))
fmt.Println(reflect.TypeOf(name2))

}