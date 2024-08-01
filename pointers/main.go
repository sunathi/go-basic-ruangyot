package main

import "fmt"

func main() {
	// b := new(int)
	// c := new(*int)  // double

	// a := 1
	// b = &a // เก็บ address ของ a เป็น value
	// c = &b // เก็บ address ของ b เป็น value
	// fmt.Printf("address a %v\n", &a) 
	// fmt.Printf("address b %v\n", &b)
	// // 
	// fmt.Printf("address c %v\n", &c)

	// fmt.Printf("value a %v\n", a)
	// fmt.Printf("value b %v\n", b)
	// fmt.Printf("value c %v\n", c)

	// fmt.Printf("*b %v\n", *b) // ดึงค่าของตัวที่ point ไป address ของ a ค่า value ของ a
	// fmt.Printf("*c %v\n", *c)
	// fmt.Printf("**c %v\n", **c) // ดู 2 ชั้น => c => b => a

	// // pass by pointers 
	m := 2
	var n *int = &m
	double(n)
	fmt.Println(*n)
	fmt.Println(m)

	// pass by reference 
	p := 2
	double(&p)
	fmt.Println(p)


}

func double(n *int) {
	*n *= 2
}
