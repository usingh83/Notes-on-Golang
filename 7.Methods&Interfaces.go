package main

import (
	"fmt"
)

type testiface interface {
	SayHello()
	Say(s string)
	Increment()
	GetInternalVal() int
}

type testConcreteImpt struct {
	i int
}

func NewTestConcreteImpt() testiface {
	return new(testConcreteImpt)
}

func NewTestConcreteImptWithV(v int) testiface {
	return &testConcreteImpt{i: v}
}

func (tst *testConcreteImpt) SayHello() {
	fmt.Println("Hello world")
}

func (tst *testConcreteImpt) Say(s string) {
	fmt.Println(s)
}

func (tst *testConcreteImpt) Increment() {
	tst.i++
}

func (tst *testConcreteImpt) GetInternalVal() int {
	return tst.i
}

type testEmbeding struct {
	*testConcreteImpt
}

func testEIface(i interface{}) {
	//fmt.Println(i)
	/*if v, ok := i.(int); ok { //type assertion
		fmt.Println("I am a int value", v)
	} else {
		fmt.Println("I am not an integer type")
	}*/
	switch val := i.(type) { //type switch
	case int:
		fmt.Println("I am a Int", val)
	case string:
		fmt.Println("I am a String", val)
	default:
		fmt.Println("I am something else", val)
	}
}

func main() {
	//var tst testiface = NewTestConcreteImpt()
	//var tst testiface = &testConcreteImpt{} //or new(testConcreteImpt)
	var tst testiface = NewTestConcreteImptWithV(5)
	tst.SayHello()
	tst.Say("Hello again!!!!!")
	tst.Increment()
	tst.Increment()
	fmt.Println(tst.GetInternalVal())
	te := testEmbeding{testConcreteImpt: &testConcreteImpt{i: 50}}
	te.SayHello()
	te.Say("Embeded")
	te.Increment()
	te.Increment()
	fmt.Println(te.GetInternalVal())
	testEIface(3)
	testEIface("Hi!! How are you!!")
	testEIface(tst)
}
