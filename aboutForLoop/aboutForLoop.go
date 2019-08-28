package aboutForLoop

//package main

import "fmt"

var aList = []int{1, 2, 3, 4, 5}
var aMap = map[string]int{"A": 1, "B": 2}

type AKindOfStruct struct {
	AttributionA int
}

var aStructList = []AKindOfStruct{{1}, {2}, {3}, {4}, {5}}

//for-loop of list
func FunctionA(aList []int) {

	for index, value := range aList {
		fmt.Println(index, value)
	}
}

//for-loop of map
func FunctionB(aMap map[string]int) {

	for key, value := range aMap {
		fmt.Println(key, value)
	}
}

func FunctionX(aMap []AKindOfStruct) {

	for key, value := range aMap {
		fmt.Println(key, value)
	}
}

func main() {
	FunctionA(aList)
	FunctionB(aMap)
	FunctionX(aStructList)
}
