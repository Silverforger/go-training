package main

import (
	"fmt"
	"strconv"
)

func boxCheck(q, a, b, c int) string {
	caseNum := strconv.Itoa(q)
	var result string
	if a <= 20 && b <= 20 && c <= 20 {
		result = "Case " + caseNum + ": good"
		return result
	} else {
		result = "Case " + caseNum + ": bad"
		return result
	}
}

func main() {
	var n, x, y, z int
	caseList := make([]string, 0)
	fmt.Scan(&n)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&x, &y, &z)
		caseList = append(caseList, boxCheck(i, x, y, z))
	}
	for i := 0; i < len(caseList); i++ {
		fmt.Println(caseList[i])
	}
}
