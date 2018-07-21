package main

import (
	"fmt"
)

// selectionSort performs selection sort, it's an in-place sort of O(n*n) complexity.
func selectionSort(array []int) {
	x := len(array)
	for i := 0; i < x; i++ {
		min := i
		for j := i; j < x; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		println("Iteration",i)
		println("Current value",array[i],"will be replaced with", array[min])
		array[i], array[min] = array[min], array[i]
	}
}

func main(){
	array := []int{7,5,3,2,4}
	fmt.Println("Unsorted -> ",array)
	selectionSort(array)
	fmt.Println("Sorted -> ",array)
}