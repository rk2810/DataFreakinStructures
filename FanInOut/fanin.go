package main

import(
	"math/rand"
	"time"
	"fmt"
)

func producer(ch chan <- int, name string) { // Write only channel
	for {
		n := rand.Intn(100)
		time.Sleep(time.Duration(n) * time.Millisecond)
		fmt.Printf("channel %s, sending %d\n", name, n)
		ch <- n
	}
}

func consumer(ch <- chan int) { // Read only channel
	for n := range ch {
		fmt.Println("<-", n)
	}
}

func fanIn(chA, chB <- chan int, chC chan <- int) {
	var n int
	for {
		select {
		case n = <- chA:
			chC <-n
		case n = <- chB:
			chC <- n
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA, "A")
	go producer(chB, "B")
	go consumer(chC)

	fanIn(chA, chB, chC)
}
