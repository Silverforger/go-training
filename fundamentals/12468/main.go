package main

import "fmt"

func pressCount(a, b int) int {
	if (b - a) < (a + 100 - b) {
		return b - a
	} else {
		return a + 100 - b
	}
}

func main() {
	var x, y int
	pressList := make([]int, 0)
	_, err := fmt.Scan(&x, &y)
	inputCounter := 1
	if x != -1 && y != -1 {
		for err == nil && x != -1 && y != -1 {
			pressList = append(pressList, pressCount(x, y))
			inputCounter++
			if inputCounter <= 200 {
				_, err = fmt.Scan(&x, &y)
			} else {
				break
			}
		}
		for i := 0; i < len(pressList); i++ {
			fmt.Println(pressList[i])
		}
	}
}
