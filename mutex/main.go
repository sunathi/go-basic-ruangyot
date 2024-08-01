package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	balance int = 1000
)

func main() {
	doneCh := make(chan bool, 3)

	go updateBalance(doneCh, 100)
	go updateBalance(doneCh, 200)
	go updateBalance(doneCh, 300)

	<-doneCh
	<-doneCh
	<-doneCh

	fmt.Println(balance)
}

func updateBalance(doneCh chan<- bool, amount int) { // ถ้าเสร็จแล้ว ใส่ true ใน doneCh
	mu.Lock()

	fmt.Println("Updating balance")
	time.Sleep(1 * time.Second)

	balance -= amount
	doneCh <- true

	mu.Unlock()
	fmt.Println("Balance Updated")
}
