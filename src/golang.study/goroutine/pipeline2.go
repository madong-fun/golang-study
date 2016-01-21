package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int) //可读写
	squares := make(chan int)
	go counter(naturals) //隐式转换
	go squarer(squares, naturals)
	printer(squares)
}

//计数
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)

}

//平方
/**
*  chan<- 只写
 */
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

//打印
/**
 * <-chan 只读
 */
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
