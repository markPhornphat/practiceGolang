package main

/*
For terminal command
- Go mod init "github path (github.com/marklileo/'name')": initial the module that used to manage Go
- Go build "fileName": Use to build binary file from source code
- Go get "dogo":Use package from other sources
		Ex: dogo (package that's used to run code automatically) --> go get github.com/liudng/dogo

*/

// "library_name" for 1 library
import "fmt" //Use as standard i/o
 
func main() {
	fmt.Println("Hello World!")
}