package main

import "fmt"

func hello() {
	fmt.Println("Hello!")
}

func sum(x int, y int) int { 
	return x+y;
}

func plus3Value(val1 int, val2 int, val3 int) int {
	return val1 + val2 + val3; 
}

func main() {
	hello()
	fmt.Println(sum(5,9))
	fmt.Println(plus3Value(9,9,9))
}