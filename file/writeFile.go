package main

import (
	"fmt"
	"os"
)

func check (err error){
	if err != nil{
		fmt.Println(err)
	}
}

func main() {
	
	//Write a file
	data1 := []byte("hello\n World") // "byte" type for input to the file
	err := os.WriteFile("file/data.txt", data1, 0644) //0644 is file type  //Write into a file NOT append 
	if err != nil{
		fmt.Println(err)
	}
	
	f, err := os.Create("employeeName")
	if err != nil{
		fmt.Println(err)
	}

	defer f.Close() //To close the file when the program finished

	data2 := []byte("Phornphat Chanthanarak");
	os.WriteFile("employeeName", data2, 0644)
}