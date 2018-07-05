package main

import (
	"math/rand"
	"fmt"
	"time"
)

func process2(ch chan int) {
	n := rand.Intn(1000) // not very random
	time.Sleep(time.Duration(n) * time.Millisecond)
	ch <- n
}

func main() {
	ch := make(chan int)
	go process2(ch)

	fmt.Println("Waiting for channel...")
	n := <- ch
	fmt.Printf("Process took %dms\n", n)
}
