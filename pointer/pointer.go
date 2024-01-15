package main

import "fmt"

func zeroPointer(iPointer *int){
	*iPointer = 0
}

func main() {
	i := 1
	zeroPointer(&i);
	fmt.Println("i from zeroVal=", i)
}