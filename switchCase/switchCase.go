package main

import "fmt"

func main() {
	input := "green"
	switch input {
	case "blue":
		fmt.Println("#000FF")
	case "green":
		fmt.Println("008000")
	case "pink":
		fmt.Println("#FFC0CB")
	default: //If the case isn't exist 
		fmt.Println("Invalid color")
	}
}