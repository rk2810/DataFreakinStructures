package main

import "fmt"

func permuation(array []int, l int, r int) { // l is left most index and r is right most index
	if (l == r) {
		fmt.Println(array)
	}
	for i := l; i <= r; i++ {
		array[l], array[i] = array[i], array[l] // swap the values at indexes
		permuation(array, l+1, r) // next index
		array[i], array[l] = array[l], array[i] // restore old config for backtrack
	}
}

func main(){
	var n int
	println("Enter number of elements in array:")
	fmt.Scanln(&n)
	println("Enter array:")
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
	}
	permuation(arr, 0, n - 1) //  array, left most index(obviously 0), right most index(length - 1)
}