package main

import (
	"fmt"
	"time"
)

func main(){
	jobCh := make(chan int, 10)
	resultCh := make(chan int, 10)

	for i := range 10 {
		jobCh <- i + 1 // start ที่ 1
	}
	close(jobCh);
	for range 2 {
		go double(jobCh, resultCh);
	}
	
	for range 10 {
		result := <-resultCh
		fmt.Println(result)
	}
}

func double(a <- chan int, b chan<- int ){
	for j := range a {
		time.Sleep(2* time.Second)
		b <- j * 2
	}
}

// worker cool 