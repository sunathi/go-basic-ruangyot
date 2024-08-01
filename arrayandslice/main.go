package main

import "fmt"

func main() {
	// array can't append
	var a [3]int = [3]int{1, 2, 3}

	fmt.Println(a)
	a[0] = 4
	fmt.Println(a)

	for i := range a {
		fmt.Println("a[i]", a[i])
	}

	for i, v := range a {
		fmt.Println("i, v", i, v)
	}

	for _, v := range a {
		fmt.Println("v", v)
	}

	p := [3]int{1, 2, 3}
	fmt.Println("in main array")
	for i := range p {
		fmt.Printf("%v\n", &p[i])
	}
	double(p)
	fmt.Println("p", p)

	// slice
	var b []int = []int{1, 2, 3}
	fmt.Println("b", b)

	c := make([]int, 0)
	c = append(c, 1, 2, 3)
	fmt.Println("c", c)
	c[2] = 5
	fmt.Println("c", c)

	m := []int{1, 2, 3}
	fmt.Println("in main slide")
	for i := range m {
		fmt.Printf("%v\n", &m[i])
	}
	double2(m)
	fmt.Println("m", m)

	// ใส่ func แล้วเปลี่ยนค่าที่ m เลย
	double3(m)
	fmt.Println("m correct func", m)

}

func double(n [3]int) {
	fmt.Println("in double array")
	for i := range n {
		fmt.Printf("%v\n", &n[i])
		n[i] *= 2
	}
}

func double2(n []int) {
	fmt.Println("in double slide")
	for i := range n {
		fmt.Printf("%v\n", &n[i])
		n[i] *= 2
	}
}

// correct func => have return => pass by reference by defalt not copy
func double3(nums []int) []int {
	newNums := make([]int, len(nums))
	fmt.Println("in double slide correct func")
	for i := range nums {
		fmt.Printf("%v\n", &nums[i])
		newNums[i] = nums[i] * 2
	}
	return newNums
}

// add array in func double not change value because it's copy then address is difference
// but slice is the same address value is changed
// so return value !!! เปลือง memory แต่จำเป็น
