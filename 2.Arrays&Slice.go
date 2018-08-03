package main

import (
	"fmt"
)

func main() {
	//var a [5]int=[5]int{1,2,3,4,5}
	a:=[5]int{1,2,3,4,5}
	myslice:=[]int{1,2,3,4,5}
	//var b [5]int
	fmt.Println(a)
	myslice=append(myslice,6)
	fmt.Println(myslice)
	s:=make([]int,5)
	s[0],s[1],s[2],s[3],s[4]=1,2,3,4,5//capacity: cap() length: len()
	fmt.Println(s)
	s1:=s[1:3]
	fmt.Println(s1)
	fmt.Println(s1[:cap(s1)])
	s2:=make([]int,2)
	copy(s2,s[1:3])
	fmt.Println(s2[:cap(s2)])
	fmt.Println(gettwo(s,2,4))
}

func gettwo(a []int,s,e int) []int{	
	//return a[s:e]  //if a has 100 elements, then the return value will have access to 100-s elements
	a1:=make([]int,e-s)
	copy(a1,a[s:e])
	return a1

}
