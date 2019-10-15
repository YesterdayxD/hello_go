//package tcpServer
package main
import (
	"fmt"
	"net"
)

//package tcpServer


func main(){
	//监听
	listener, err :=net.Listen("tcp","127.0.0.1:8000")
	if err != nil{
		fmt.Println("err=",err)
		return
	}
	defer listener.Close()
	//阻塞等待用户连接
	conn,err := listener.Accept() //conn 得到了客户端的数据
	if err!=nil {
		fmt.Println("err=",err)
		return
	}
	buf:=make([]byte,1024)
	n,err := conn.Read(buf) //将conn的数据读取并给buf 同时返回一个数据长度n
	if err != nil {
		fmt.Println("err=",err)
		return
	}
	fmt.Println("buf=",string(buf[:n]))
	defer conn.Close()
}

