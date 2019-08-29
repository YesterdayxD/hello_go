package aboutClosure
//package main

import "fmt"

//AClosureFunction的返回值是一个函数
//func(int) int这又可以看做是一个函数，这个函数的返回值是一个int型
func AClosureFunction() func(int) int {
	return func(i int) int {
		return i
	}
}
func AnotherClosureFunction() func() {
	n := 0
	return func() {
		fmt.Println(n)
		n += 1
	}
}
func ThirdClosureFunction(str string) func(i int) int {
	fmt.Println(str)
	n:=0
	return func(i int)int {
		fmt.Println(n)
		n += 1
		return i
	}
}
func main() {
	aFunc := AnotherClosureFunction()
	aFunc()
	aFunc()
	aFunc()
	bFunc:=ThirdClosureFunction("hello")
	fmt.Println(bFunc(100))
	fmt.Println(bFunc(200))
	fmt.Println(bFunc(300))
}
