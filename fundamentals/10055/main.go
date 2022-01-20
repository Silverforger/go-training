package main

import "fmt"

func main() {
	var x, y int
	resultList := make([]int, 0)
	_, err := fmt.Scan(&x, &y)
	for err == nil {
		resultList = append(resultList, y-x)
		_, err = fmt.Scan(&x, &y)
	}
	if err != nil {
		for i := 0; i < len(resultList); i++ {
			fmt.Println(resultList[i])
		}
	}
}
