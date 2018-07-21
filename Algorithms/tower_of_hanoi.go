package main

import "fmt"

func towerOfHanoi(height int, x string, z string, y string) {
	if (height == 1) {
		println("Move disk ", height, "from ", x, "to", z)
	} else {
		towerOfHanoi(height-1, x, y, z)
		println("Move disk ", height,"from ", x, "to", z)
		towerOfHanoi(height-1, y, z, x)
		towerOfHanoi(height-1, y, z, x)
	}
}

func main() {
	fmt.Println("Enter the number of disks :")
	var n int
	fmt.Scanln(&n) // be sane with this tho, 2^n-1 iterations
	// Since there are 3 towers
	x := "Tower A" // Source tower
	y := "Tower B" // Aux tower
	z := "Tower C" // Destination tower
	towerOfHanoi(n,x,z,y)
}	