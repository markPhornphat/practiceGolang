package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("file/accessories.csv") //Open the file 
	
	if err != nil {
		panic(err) //print the err
	}

	scanner := bufio.NewScanner(file) //To read value in file

	count := 1
	for scanner.Scan(){ //Loop for every row
		line := scanner.Text() //Retrieve information
		item := strings.Split(line, ",") //Split commas(,) for get a specific attribute
		fmt.Printf("Line %d : %s\n", count, item[4])
		count ++
	}
}