package main

import "fmt"

func triCheck(a, b, c int) string {
	if a <= 30000 && b <= 30000 && c <= 30000 {
		if (a*a+b*b == c*c) || (a*a+c*c == b*b) || (b*b+c*c == a*a) {
			return "right"
		} else {
			return "wrong"
		}
	} else {
		return "Input exceeded 30000"
	}
}

func main() {
	var x, y, z int
	resultList := make([]string, 0)
	_, err := fmt.Scan(&x, &y, &z)
	for err == nil {
		if x != 0 && y != 0 && z != 0 {
			result := triCheck(x, y, z)
			resultList = append(resultList, result)
			_, err = fmt.Scan(&x, &y, &z)
		} else {
			break
		}
	}
	for i := 0; i < len(resultList); i++ {
		fmt.Println(resultList[i])
	}
}
