package main

import "fmt"

var productName [4]string

func main() {
	productName[0] = "Apple";
	productName[1] = "Macbook";
	productName[2] = "iPhone";
	productName[3] = "iPad";

	price := [4]float64{1000,2000,3000,4000}
	fmt.Println(productName); // [   ]
	fmt.Println(price) // [0, 0, 0, 0]
}