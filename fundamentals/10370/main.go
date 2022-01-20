package main

import "fmt"

func aboveAve(grades []int) float32 {
	var total, numAAve int
	var average, percentAAve float32
	for _, v := range grades {
		total += v
	}
	average = float32(total) / float32(len(grades))
	for _, p := range grades {
		if float32(p) > average {
			numAAve++
		}
	}
	percentAAve = (float32(numAAve) / float32(len(grades))) * 100
	return percentAAve
}

func main() {
	var n1, n2, x int
	class := make([]int, 0)
	percentList := make([]float32, 0)
	fmt.Scan(&n1)
	for i := 0; i < n1; i++ {
		fmt.Scan(&n2)
		for j := 0; j < n2; j++ {
			fmt.Scan(&x)
			class = append(class, x)
		}
		percentList = append(percentList, aboveAve(class))
		class = class[:0]
	}
	for i := 0; i < len(percentList); i++ {
		fmt.Println(fmt.Sprintf("%.3f", percentList[i]))
	}
}
