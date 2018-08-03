package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Printf("Hello ")
	fmt.Printf("world")
	s := struct {
		i int
		f float32
	}{i: 3, f: 3.3}
	fmt.Printf("%v \n", s)
	fmt.Printf("i is %d and f is %f ", s.i, s.f)
}
