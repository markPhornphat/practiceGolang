package main

import "fmt"

func main() {
	score := 0;

	fmt.Print("Enter your score: ");
	fmt.Scan(&score);

	if score >= 80 {
		fmt.Println("A");
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else if score >= 50 {
		fmt.Println("D")
	} else if score < 50 {
		fmt.Println("F")
	}
}