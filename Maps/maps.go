package main

import "fmt"

var product = make(map[string]float64) //Initialize Maps type

func main() {
	fmt.Println(product)
	
	//Add value
	product["Macbook"] = 40000
	product["iPhone"] = 30000
	product["iPad"] = 10000
	fmt.Println("products = ",product)

	//Delete value
	delete(product, "iPad")
	fmt.Println("products after delete = ",product)
	

	//Update value
	product["iPhone"] = 39900
	fmt.Println("products after update = ",product)

	//To access the data in Maps
	iPhonePrice := product["iPhone"]
	fmt.Println("iPhone price = ", iPhonePrice)
	
	//Initialize the data
	courseName := map[string]string{"101" : "Java" , "102":"Python"}
	fmt.Println("course name = ", courseName)
}