package partyrobot

import (
	"fmt"
	"math"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// tableNo := strconv.Itoa(table)
	// switch len(tableNo) {
	// case 1:
	// 	tableNo = "00" + tableNo
	// case 2:
	// 	tableNo = "0" + tableNo
	// }
	roundedDis := math.Round(distance*10) / 10
	finalMessage := fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, table, direction, roundedDis, neighbor)
	return finalMessage
}
