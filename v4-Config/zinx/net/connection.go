package net

import (
	"Zinx/Project/Zinx/v4-Config/zinx/iface"
	"fmt"
	"net"
)

//封装原生tcpconn
type Connection struct {
	TCPconn  *net.TCPConn
	Connid   uint32
	IsClosed bool
	//Callback iface.CallBack
	router iface.IRouter
}

//创建一个Connection
func Newconn(conn *net.TCPConn, id uint32, router iface.IRouter) iface.Iconnection {
	return &Connection{
		TCPconn:  conn,
		Connid:   id,
		//Callback: callback,
		router:router,
	}
}

//返回原生tcpconn
func (c *Connection) GetTcpConn() *net.TCPConn {
	return c.TCPconn
}

//返回连接id
func (c *Connection) GetConnId() uint32 {
	return c.Connid
}

//封装读取数据，和返回数据功能
func (c *Connection) Start() {
	fmt.Println("Conn start... id:", c.Connid)
	buf := make([]byte, 4096)
	for {
		n, err := c.TCPconn.Read(buf)
		if err != nil {
			fmt.Println("Read err ", err)
			return
		}
		data := string(buf[:n])
		fmt.Println("客户端发送过来数据")
		req := NewRequest(c, []byte(data), uint32(n))
		//转成大写，返回给客户端
		//c.Callback(req)
		c.router.Handle(req)
	}
}

func (c *Connection) Stop() {
	fmt.Println("Conn stop... id:", c.Connid)
	if !c.IsClosed {
		fmt.Println("Conn已关闭")
		return
	}

	_ = c.TCPconn.Close()
	c.IsClosed = true
}

//发送数据返回
func (c *Connection) Send(DATA string) {
	_, err := c.TCPconn.Write([]byte(DATA))
	if err != nil {
		fmt.Println("Write err", err)
		return
	}
	fmt.Println("数据已成功返回")
}
