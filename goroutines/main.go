package main

import (
	"fmt"
	"sync"
)

func main() {
	// fmt.Println("Hello 1")
	// time.Sleep(1 * time.Second)

	// go hello() //แยกไปทำงาน
	fmt.Println("Hello 2")
	// time.Sleep(1 * time.Second)
	var wg sync.WaitGroup // ทำตัวไหนก็ได้
	wg.Add(3)
	for i := range 3 {
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done() // &wg
			fmt.Println("Hello another thread")
			fmt.Println(i)
		}(i, &wg)
	}
	wg.Wait()
}

// func hello() {
// 	for {
// 		fmt.Println("Hello another")
// 	}
// }
