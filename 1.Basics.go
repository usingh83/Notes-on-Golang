package main

import (
	"fmt"
)

//var x uint8 = 2

func main() {
	fmt.Println(CompareNum(1, 2))
	//fmt.Println(x)
	/*switch x{
	case 1:
		fmt.Println("x is one")
	case 2:
		fmt.Println("x is two")
		fallthrough
	case 3:
		fmt.Println("x is three")
	}*/
	switch x := 2; x {
	case 1:
		fmt.Println("x is one")
	case 2:
		fmt.Println("x is two")
	case 3:
		fmt.Println("x is three")
	}
	/*for i:=0;i<=10;i++ {
		fmt.Println(i)
	}*/
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
func CompareNum(i1, i2 int) (bool, int) {
	/*if i1>i2 {
		fmt.Println("first number is greater")
		return false,i1-i2
	} else if i2>i1 {
		fmt.Println("Second number is greater")
		return false,i2-i1
	}
	fmt.Println("both numbers are equal")
	return true,0
	*/
	switch {
	case i1 > i2:
		fmt.Println("first number is greater")
		return false, i1 - i2
	case i2 > i1:
		fmt.Println("Second number is greater")
		return false, i2 - i1
	}
	fmt.Println("both numbers are equal")
	return true, 0

}
