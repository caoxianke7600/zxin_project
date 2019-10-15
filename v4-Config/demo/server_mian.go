package main

import (
	"Zinx/Project/Zinx/v4-Config/zinx/iface"
	"Zinx/Project/Zinx/v4-Config/zinx/net"
	"fmt"
	"strings"
)

type TestRouter struct {  //继承Router结构体，可以继承Handle方法
	net.Request
}

func (tr *TestRouter)Handle(req iface.IRequest)  {
	fmt.Println("TestRouter called")
	data := req.GetData() //获取Request的接口绑定的方法
	Conn := req.GetConn()
	fmt.Println("Userbussiness called,data:", string(data))
	DATA := strings.ToUpper(string(data)) //转大写
	//调用send发送数据
	Conn.Send(DATA)
}

func main() {
	server := net.NewServer("Zinxv1.0")

	//注册路由
	server.AddRouter(&TestRouter{})
	server.Server()
}
