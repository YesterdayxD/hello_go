package aboutInterface

//定义一种接口类型
type AKindOfInterface interface {
	FunctionA() string
	FunctionB() string
}

//定义一个结构体
type AKindOfStrut struct {
	AttributionA int
	AttributionB string
}

//在结构体AKindOfStrut实现AKindOfInterface借口的函数（所有的函数）
func (s *AKindOfStrut) FunctionA() string {
	//fmt.Println("hello")
	return "hello,A"
}

func (s *AKindOfStrut) FunctionB() string {
	//fmt.Println("hello")
	return "hello,B"
}
//通过以下语法判断某个数据类型是否全部了所有的借口
//如果没有实现，则会编译失败
// 	type T struct{}
//  var _ I = &T{}
//其中 T 为结构体（数据类型），I为接口
var _ AKindOfInterface = &AKindOfStrut{}

func main() {
	//通过以下语法判断某个数据类型是否全部了所有的借口
	//如果没有实现，则会编译失败
	// 	type T struct{}
	//  var _ I = &T{}
	//其中 T 为结构体（数据类型），I为接口
	var _ AKindOfInterface = &AKindOfStrut{}
}
