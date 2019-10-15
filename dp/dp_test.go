package main

import (
	"testing"
	//"math/rand"
)

func TestRob(t *testing.T) {

	for index,unit :=range[]struct{
		house []int
		money int
	}{
		{[]int{1,2,5,2,1,3},9},
		{[]int{1},1},
		{[]int{1,2},2},
		{[]int{1,2,3},4},
		{[]int{},0},
	}{
		if unit.money !=Rob(unit.house){
			t.Errorf("wrong index %d",index)
		}
	}

	//house:=[]int{1,2,5,2,1,3}
	//fmt.Println(house)
	//if 9==Rob(house){
	//	t.Log("right")
	//	//t.Error("right")
	//}else{
	//	t.Error("wrong")
	//}
}

func BenchmarkRob(b *testing.B) {
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		_=Rob([]int{1,2,5,2,1,3})
	}
}
func BenchmarkJump(b *testing.B) {
	for i:=0;i<b.N;i++{
		_=Jump(i)
	}
}
