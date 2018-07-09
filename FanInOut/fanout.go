package main

import(
	"math/rand"
	"time"
	"fmt"
)

func producer(ch chan <- int) { // Write only channel
	for {
		n := rand.Intn(100)
		time.Sleep(time.Duration(n) * time.Millisecond)
		fmt.Println("->", n)
		ch <- n
	}
}

func consumer(ch <- chan int, name string) { // Read only channel
	for n := range ch {
		fmt.Println("Consumer", name, "<-", n)
	}
}

func fanOut(chA <- chan int, chB, chC chan <- int){
	for n := range chA{
		if n <= 50{
			chB <- n
		} else {
			chC <- n
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA)
	go consumer(chB, "B")
	go consumer(chC, "C")

	fanOut(chA, chB, chC)
}
