package main

import "fmt"

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("Result: ", result)
}

func loop(){
	for i := 0; i < 10; i++ {
		fmt.Println("i = ",i);
	}
}

func deferLoop(){
	for j := 0; j < 10; j++ {
		defer fmt.Println("j = ", j)
	}
}

func main() {
// 	fmt.Println("Start Calculator")
// 	defer fmt.Println("End")
// 	defer add(10,20);
// 	defer add(15,23);
// 	defer add(1,21);
	loop()
	deferLoop()
}