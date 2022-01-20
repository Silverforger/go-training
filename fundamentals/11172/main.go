package main

import "fmt"

func main() {
	var n, x, y int
	var operator string
	fmt.Scan(&n)
	operatorList := make([]string, 0)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&x, &y)
		if x < y {
			operator = "<"
		} else if x == y {
			operator = "="
		} else if x > y {
			operator = ">"
		}
		operatorList = append(operatorList, operator)
	}
	for i := 0; i < len(operatorList)+1; i++ {
		fmt.Println(operatorList[i])
	}
}
