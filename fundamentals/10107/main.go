package main

import (
	"fmt"
	"sort"
)

func medianChecker(array []int) int {
	var median int
	if len(array)%2 != 0 {
		median = array[len(array)/2]
		return median
	} else if len(array)%2 == 0 {
		median = (array[len(array)/2-1] + array[len(array)/2]) / 2
		return median
	}
	return median
}

func main() {
	var x, intCounter int
	inputList := make([]int, 0)
	_, err := fmt.Scan(&x)
	intCounter++
	for err == nil {
		inputList = append(inputList, x)
		sort.Ints(inputList[:])
		if intCounter <= 10000 {
			fmt.Println(medianChecker(inputList))
			_, err = fmt.Scan(&x)
			intCounter++
		} else {
			break
		}
	}
}
