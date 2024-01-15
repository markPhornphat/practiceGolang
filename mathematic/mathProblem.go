package main

import "fmt"

var operator[] string
var result float64
var input []int

func getNum(i int){
	temp := 0;
	fmt.Printf("Number %d :",i+1);
	fmt.Scan("%d", temp);
	input = append(input, temp);

	getOperator(i);
}

func getOperator(i int){
	temp := "+"

	fmt.Println("Operator : ")
	fmt.Scanln("%d", temp)
	operator = append(operator, temp)
}

func calculate(i int){
	for j := 0; j < i; j++ {
		// result = j + i;
	}
}

func main() {
	count := 0
	fmt.Printf("How much number would you like to calculate: ")
	fmt.Scanf("%d", &count);

	for i := 0; i < count; i++ {
		getNum(i);
		if(i == count - 1){
			calculate(i);
			break;
		}
	}
}