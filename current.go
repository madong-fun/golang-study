package main

import (
	"fmt"
	"time"
)

var c chan int

func main() {
	c = make(chan int)

	go ready("Tea", 2)
	go ready("Coffee", 1)

	fmt.Println("waiting!!!")

	fmt.Println(<-c)
	fmt.Println(<-c)

}

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Print("ready...")
	c <- 1
}
