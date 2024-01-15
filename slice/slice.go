package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Python"}
	courseName = append(courseName, "Golang", "C", "HTML")	
	
	courseWeb := courseName[2:5]; //Choose only the exclude index number 2 to 5 
	fmt.Println(courseWeb);

	courseWeb = courseName[:2]
	fmt.Println(courseWeb);
}