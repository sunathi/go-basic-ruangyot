package main

import "fmt"

func main() {
	// reuse
	hello("John")
	hello("Marry")

	fmt.Println("add", add(4, 2))

	result := sub(4, 2)
	fmt.Println("result sub", result)

	fmt.Println(div(4, 0))

	// arenemos func ไม่จำเป็นไม่ต้องใช้
	name := "soi"
	func(name string) {
		fmt.Println("Hello", name)
	}(name)

	// *** be careful !!! *** value chaged
	lastname := "kk"
	func() {
		lastname = "jj"
		fmt.Println("Hello", lastname)
	}()

	result2 := func(c, d int) int {
		return c + d
	}(1, 2)
	fmt.Println(result2)
}

// func hello(name, lastname string) string {
// 	return ""
// }

func hello(name string) {
	fmt.Println("hello", name)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func div(a, b int) int {
	if b == 0 {
		fmt.Println("Error: Division by 0")
		return 0
	}
	return a / b
}
