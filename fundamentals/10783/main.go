package main

import "fmt"

func main() {
	var n, x, y, oddsum int
	fmt.Scan(&n)
	oddsums := make([]int, 0)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&x, &y)
		for i := x; i < y+1; i++ {
			if (i % 2) != 0 {
				oddsum += i
			}
		}
		oddsums = append(oddsums, oddsum)
		oddsum = 0
	}
	for i := 0; i < len(oddsums)+1; i++ {
		fmt.Println(oddsums[i])
	}
}
