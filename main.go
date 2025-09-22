package main

import "fmt"

func main() {
	myMap := make(map[string]string)

	myMap["name"] = "John"
	myMap["names"] = "Johné"
	myMap["namess"] = "John2"
	myMap["namesd"] = "John3"
	myMap["namesf"] = "John4"

	for key, values := range myMap {
		fmt.Println(key, values)
	}
}
