package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println("i = ", i)
		i += 1
	}

	const count = 15;
	for j := 0; j < count; j++ {
		fmt.Println("j = ", j)
		
	}

}