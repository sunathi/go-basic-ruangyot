package main

import (
	"fmt"
	"reflect"
)

func main() {
	// implicit type: fixed type
	var a int = 10
	fmt.Println(reflect.TypeOf(a), a)
	// dynamic type: find type automatically
	b := 20
	c := "hello"
	fmt.Println(reflect.TypeOf(b), b)
	fmt.Println(reflect.TypeOf(c), c)

}
