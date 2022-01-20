package main

import "fmt"

func main() {
	var n, x, y, z, survivor int
	fmt.Scan(&n)
	survivorList := make([]int, 0)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&x, &y, &z)
		if (x < y && x > z) || (x < z && x > y) {
			survivor = x
		} else if (y < x && y > z) || (y < z && y > x) {
			survivor = y
		} else if (z < y && z > x) || (z < x && z > y) {
			survivor = z
		}
		survivorList = append(survivorList, survivor)
	}
	for i := 0; i < len(survivorList); i++ {
		fmt.Println(survivorList[i])
	}
}
