package main

import "fmt"

var graph = map[int][]int{
	1: {2,3},
	2: {4,5},
	3: {6},
}
func main() {
	// a := map[int]int{
	// 	1: 1,
	// 	2: 6,
	// 	3: 3,
	// }
	// fmt.Println(a[2])
	// if _, ok := a[100]; ok {
	// 	fmt.Println("OK")
	// } else {
	// 	fmt.Println("NOT OK")
	// }

	// "make"
	// b := make(map[int]int)
	// b[1] = 1
	// fmt.Println(b)
	dfs(graph, 1, make(map[int]bool))

}

// HashMap ใช้ memory มากกว่า for ด้วย slice แต่เร็วกว่า

func dfs(graph map[int][]int, vertex int, visited map[int]bool) {
	visited[vertex] = true
	for _, v := range graph[vertex] {
		fmt.Printf("->%d", v)
		dfs(graph, v, visited)
	}
}