package iface

import "net"

type Iconnection interface {
	GetTcpConn() *net.TCPConn
	GetConnId() uint32
	Start()
	Stop()
	Send(DATA string)
}

//定义一个回调函数
type CallBack func(request IRequest) //调到server.go中的func(Conn net2.Connection, data []byte)类型函数
