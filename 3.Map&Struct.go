package main

import (
	"fmt"
)

type person struct{
	name string
	age int
	address string
}

func main() {
	//map
	//var m map[string]int=make(map[string]int)
	/*m:=make(map[string]int)
	m["first"]=1
	m["Second"]=2
	*/
	m:=map[string]int{
		"first":1,
		"second":2,
	}
	fmt.Println(m)
	fmt.Println(m["first"])
	//_,ok:=m["third"]
	//fmt.Println(ok)
	if v,ok:=m["first"];ok{
		fmt.Println(v)
	}
	delete(m,"first")
	fmt.Println(m)
	
	//structs
	
	jason:=person{
		name:"json b",
		age:33,
		address:"home",
	}
	fmt.Println(jason)
	fmt.Println(jason.name)
		
}
