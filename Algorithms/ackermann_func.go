package main

import "fmt"
var count = 0

func ackermann(m int, n int) int{
	count++
	if m == 0{
		return n+1
	}
	if (n == 0){
		return ackermann(m-1, 1)
	}
	return ackermann(m-1, ackermann(m, n-1))
}

func main() {
	fmt.Println(ackermann(3, 2))
	fmt.Println(count)
}
