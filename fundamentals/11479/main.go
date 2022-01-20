package main

import "fmt"

func main() {
	var n, x, y, z int
	var ttype string
	fmt.Scan(&n)
	typeArray := make([]string, 0)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&x, &y, &z)
		if x+y < z || x+z < y || y+z < x {
			ttype = "Invalid"
		} else if x != y && y != z {
			ttype = "Scalene"
		} else if x == y && y == z {
			ttype = "Equilateral"
		} else if x == y || y == z || x == z {
			ttype = "Isosceles"
		}
		typeArray = append(typeArray, ttype)
	}
	for i := 0; i < len(typeArray)+1; i++ {
		fmt.Println(typeArray[i])
	}
}
