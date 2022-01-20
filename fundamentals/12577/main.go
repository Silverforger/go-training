package main

import "fmt"

func main() {
	var x string
	resultList := make([]string, 0)
	_, err := fmt.Scan(&x)
	for err == nil && x != "*" {
		if x == "Hajj" {
			resultList = append(resultList, "Hajj-e-Akbar")
			_, err = fmt.Scan(&x)
		} else if x == "Umrah" {
			resultList = append(resultList, "Hajj-e-Asghar")
			_, err = fmt.Scan(&x)
		}
	}
	for i := 0; i < len(resultList)+1; i++ {
		fmt.Println("Case ", i+1, ": ", resultList[i])
	}
}
