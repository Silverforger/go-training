package main

import "fmt"

func main() {
	var x, y int
	_, err := fmt.Scan(&x, &y)
	for err == nil {
		if -100 <= x && x <= 100 && 0 <= y && y <= 200 {
			fmt.Println(x * y * 2)
			_, err = fmt.Scan(&x, &y)
		} else if x < -100 || x > 100 {
			println("Velocity should be between -100 and 100.")
			break
		} else if y < 0 || x > 200 {
			println("Time should be between 0 and 200.")
			break
		}
	}
}
