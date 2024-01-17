package main

import (
	"fmt"
	"time"
)

func process1(c chan string, data string) { //Use "chan" to indicate chanel ==> It is a channel that used to communicate between gorutine1 and gorutine2
	c <- data //put data variable into channel c
}

func main() {
	ch := make(chan string)
	go process1(ch, "Data1")
	fmt.Println(<-ch)
	time.Sleep(5*time.Second)
}