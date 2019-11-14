package main

import (
	"fmt"
)

type Tester interface {
	Display()
	DisplayAppend(string)
	DisplayAppend2(string) string
}

type Tester2 interface {
	DisplayAppend(string)
}

type Test struct {
	s string
}

func (t *Test) Display() {
	fmt.Printf("Display:%p, %#v\n", t, t)
}

func (t Test) DisplayAppend(s string) {
	t.s += s
	fmt.Printf("DisplayAppend:%p, %#v\n", &t, t)
}

func (t *Test) DisplayAppend2(s string) string {
	t.s += s
	fmt.Printf("DisplayAppend2:%p, %#v\n", t, t)
	return t.s
}

func TestInterface(t Tester) {
	t.Display()
	t.DisplayAppend(" TestInterface")
	t.DisplayAppend2(" TestInterface")
}

func TestInterface2(t Tester2) {
	t.DisplayAppend("TestInterface2")
}

func main() {
	var test Test
	test.s = "aaa"
	fmt.Printf("%p\n", &test)
	test.Display()
	test.DisplayAppend(" raw")

	TestInterface(&test)
	//TestInterface(test) //cannot use test (type Test) as type Tester in argument to TestInterface:Test does not implement Tester (Display method has pointer receiver)

	TestInterface2(&test)
	TestInterface2(test)
}
