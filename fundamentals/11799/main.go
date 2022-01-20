package main

import (
	"fmt"
)

func max(speeds []int) int {
	maxSpeed := speeds[0]
	for _, v := range speeds {
		if v > maxSpeed {
			maxSpeed = v
		}
	}
	return maxSpeed
}

func main() {
	var n1, n2, x int
	clownSpeeds := make([]int, 0)
	speedsList := make([]int, 0)
	fmt.Scan(&n1)
	for i := 0; i < n1; i++ {
		fmt.Scan(&n2)
		for j := 0; j < n2; j++ {
			fmt.Scan(&x)
			speedsList = append(speedsList, x)
		}
		clownSpeeds = append(clownSpeeds, max(speedsList))
		speedsList = speedsList[:0]
	}
	for i := 0; i < len(clownSpeeds); i++ {
		fmt.Println("Case ", i+1, ": ", clownSpeeds[i])
	}
}
