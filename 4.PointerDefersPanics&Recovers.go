package main

import (
	"fmt"
)

func main() {
	//var pI *int
	var I int=3
	//pI=&I
	//increment(pI)
	//increment(&I)
	increment(I)
	fmt.Println(I)
	//defer fmt.Println("world 1")
	//defer fmt.Println("world 2")
	fmt.Println("Hello")
	testpanic()
	fmt.Println("world")
}
func testpanic(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println(err)
			fmt.Println("we recovered from an error")
		}
	}()
	panic("A panic happen")
}
/*
func increment(pI *int){
	*pI++
}
*/
func increment(I int){
	I++
}