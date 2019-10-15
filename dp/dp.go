package main
//package dp

import (
	"fmt"
)
func max(x,y int) int{
	if x>y{
		return x
	}else{
		return y
	}
}

//
func Rob(house []int)int{
	n:=len(house)
	if n<=0{
		return 0
	}
	if n==1{
		return house[0]
	}
	r:=make([]int,n)
	nr:=make([]int,n)
	//init
	r[0]=house[0]
	//begin
	for i:=int(1);i<n;i++{
		r[i]=nr[i-1]+house[i]
		nr[i]= max(r[i-1], nr[i-1])
	}
	return max(r[n-1],nr[n-1])
}
func Jump(n int)int{
	if n<=0{
		return 0
	}
	if n==1{
		return 1
	}
	if n==2{
		return 2
	}




	result:=0
	a:=1
	b:=2
	//ptr_result:=&result
	//ptr_a:=&a
	//ptr_b:=&b
	//for i:=3; i<=n;i++{
	//	*ptr_result=*ptr_a+*ptr_b
	//	*ptr_a=*ptr_b
	//	*ptr_b=*ptr_result
	//}
	for i:=3; i<=n;i++{
		result=a+b
		a=b
		b=result

	}

	//return *ptr_result
	return result
}

func main(){

	 //h :=[]int{1,2,5,2,1,3}
	//n:=len(h)
	//fmt.Println(n)
	//r :=make([]int,n)
	//nr :=make([]int,n)
	////init
	//r[0]=h[0]
	//fmt.Println(r)
	//fmt.Println(nr)
	//// begin
	//for i:=int(1);i<n ;i++  {
	//	r[i]=nr[i-1]+h[i]
	//	nr[i]= max(r[i-1], nr[i-1])
	//}
	//fmt.Println(r)
	//fmt.Println(nr)
	//fmt.Println(max(r[n-1],nr[n-1]))
	house:=[]int{1,2,5,2,1,3}
	money:= Rob(house)
	fmt.Println(money)

	// n
	// method 1 or 2
	// n 层数
	// 到n-1层 有m1种方法 再到n层 还是自由


	fmt.Println(Jump(0))
	fmt.Println(Jump(1))
	fmt.Println(Jump(2))
	fmt.Println(Jump(3))
	fmt.Println(Jump(4))
	fmt.Println(Jump(5))

}
