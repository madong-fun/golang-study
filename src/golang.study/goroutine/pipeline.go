package main

import (
	"fmt"
)

// pipeline  counter -> squarer ->  Printer
func main() {

	natruals := make(chan int)
	squares := make(chan int)
	//counter
	go func() {
		for i := 0; i < 100; i++ {
			natruals <- i
		}
		close(natruals)
	}()

	//squarer

	go func() {
		for x := range natruals {

			squares <- x * x
		}
		close(squares)
	}()
	//Printer
	for x := range squares {
		fmt.Println(x)
	}
}
