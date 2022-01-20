package main

import "fmt"

func main() {
	var input string
	langList := make([]string, 0)
	languages := map[string]string{
		"HELLO":        "ENGLISH",
		"HOLA":         "SPANISH",
		"HALLO":        "GERMAN",
		"BONJOUR":      "FRENCH",
		"CIAO":         "ITALIAN",
		"ZDRAVSTVUJTE": "RUSSIAN",
	}
	_, err := fmt.Scan(&input)
	for err == nil {
		if input != "#" {
			if value, ok := languages[input]; ok {
				langList = append(langList, value)
				_, err = fmt.Scan(&input)
			} else {
				langList = append(langList, "UNKNOWN")
				_, err = fmt.Scan(&input)
			}
		} else {
			break
		}
	}
	for i := 0; i < len(langList); i++ {
		fmt.Println(langList[i])
	}
}
