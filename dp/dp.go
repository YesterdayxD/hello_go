package dp

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

//
func Rob(house []int) int {
	n := len(house)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return house[0]
	}
	r := make([]int, n)
	nr := make([]int, n)
	//init
	r[0] = house[0]
	//begin
	for i := int(1); i < n; i++ {
		r[i] = nr[i-1] + house[i]
		nr[i] = max(r[i-1], nr[i-1])
	}
	return max(r[n-1], nr[n-1])
}
func Jump(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	result := 0
	a := 1
	b := 2
	//ptr_result:=&result
	//ptr_a:=&a
	//ptr_b:=&b
	//for i:=3; i<=n;i++{
	//	*ptr_result=*ptr_a+*ptr_b
	//	*ptr_a=*ptr_b
	//	*ptr_b=*ptr_result
	//}
	for i := 3; i <= n; i++ {
		result = a + b
		a = b
		b = result

	}

	//return *ptr_result
	return result
}
func Tri(n int, matrix [][]int) int {
	//recordMatrixRow := make([]int, 5)
	recordMatrix := make([][5]int, 5)
	//init
	recordMatrix[0][0] = matrix[0][0]
	//begin
	for i := 1; i < n; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 {
				recordMatrix[i][j] = recordMatrix[i-1][j] + matrix[i][j]
				fmt.Println(recordMatrix)
			} else if i == j {
				recordMatrix[i][j] = recordMatrix[i-1][j-1] + matrix[i][j]
				fmt.Println(recordMatrix)
			} else {
				recordMatrix[i][j] = max(recordMatrix[i-1][j]+matrix[i][j], recordMatrix[i-1][j-1]+matrix[i][j])
				fmt.Println(recordMatrix)

			}
		}

	}

	fmt.Println(recordMatrix)
	maxValue := recordMatrix[n-1][0]
	for i := 1; i < n; i++ {
		if recordMatrix[n-1][i] > maxValue {
			maxValue = recordMatrix[n-1][i]
		}
	}
	return maxValue
}

func main() {

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
	house := []int{1, 2, 5, 2, 1, 3}
	money := Rob(house)
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

	matrix := [][]int{
		{7},
		{8, 3},
		{8, 1, 0},
		{4, 3, 3, 6},
		{4, 4, 4, 4, 4},
	}
	fmt.Println(Tri(5, matrix))

}
