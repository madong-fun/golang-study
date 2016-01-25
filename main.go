package main

import (
	"fmt"
	"golang.study/thumbnail"
	"time"
)

func main() {
	//fmt.Println(time.Unix(1389058332, 0).Format("2006-01-02 15:04:05"))
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	go spinner(100 * time.Millisecond)
	_, err := thumbnail.ImageFile("")
	if err != nil {
		fmt.Println(err)
	}
	const n = 45
	fibN := fib(45)
	fmt.Println("\rFibonacci(%d) = %d\n", n, fibN)
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {

			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
