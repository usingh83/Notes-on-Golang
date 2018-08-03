package main

import (
	"fmt"
	"time"
)

func main() {
	qs := make(chan bool)
	//go SayHelloFromGorotinue(qs)
	go func() {
		fmt.Println("Hello from new gorotinue")
		qs <- true
	}()
	fmt.Println("Hello from main")
	v := <-qs
	fmt.Println(v)
	ic := make(chan int)
	go Periodic(ic)
	for i := range ic {
		fmt.Println(i)
	}
	//ic<-3
	_, ok := <-ic
	fmt.Println(ok)

}
func Periodic(ic chan int) {
	i := 0
	for i <= 10 {
		ic <- i
		i++
		time.Sleep(1 * time.Second)
	}
	close(ic)
}
func SayHelloFromGorotinue(qs chan bool) {
	fmt.Println("Hello from new gorotinue")
	qs <- true
}
