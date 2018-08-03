package main

import (
	"fmt"
	"time"
)

func main() {
	//buffch:=make(chan int,5)
	buffch := make(chan int, 2)
	buffch <- 3
	buffch <- 2
	//buffch<-4
	fmt.Println(<-buffch)
	fmt.Println(<-buffch)
	//fmt.Println(<-buffch)
	fmt.Println(<-waitAndSend(3, 1))
	ic := make(chan int)
	select {
	case v1 := <-waitAndSend(3, 1):
		fmt.Println(v1)
	case v2 := <-waitAndSend(5, 1):
		fmt.Println(v2)
	case ic <- 4:
		fmt.Println(<-ic)
	default:
		fmt.Println("in Default")
	}
}
func waitAndSend(v, i int) chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Duration(i) * time.Second)
		ch <- v
	}()
	return ch
}
