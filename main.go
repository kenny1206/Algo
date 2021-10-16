package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int
	go func() {
		a := <-c
		fmt.Println(a)
	}()
	c = make(chan int)
	c <- 1
	time.Sleep(time.Second)
}

