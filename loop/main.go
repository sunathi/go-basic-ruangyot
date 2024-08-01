package main

import "fmt"

func main() {
	// for i := 1; i <= 3; i++ {
	// 	fmt.Println(i)
	// }

	// var sum int
	// for i := 1; i <= 3; i++ {
	// 	sum += i
	// }

	// [0,1,2]
	for i := range 3 {
		fmt.Println(i)
	}
}
