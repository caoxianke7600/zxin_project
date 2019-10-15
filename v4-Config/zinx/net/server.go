package net

import (
	"Zinx/Project/Zinx/v4-Config/zinx/config"
	"Zinx/Project/Zinx/v4-Config/zinx/iface"
	"fmt"
	"net"
)

//定义一个server结构
type Server struct {
	IP         string
	Port       uint32
	Name       string
	TCPVersion string
	//路由字段
	router iface.IRouter
}

//创建Server方法
func NewServer(name string) iface.Iserver { //相当于多态
	return &Server{
		IP:         config.GlobalConfig.IP,
		Port:       config.GlobalConfig.Port,
		Name:       config.GlobalConfig.Name,
		TCPVersion: config.GlobalConfig.TCPVersion, //tcp,tcp4,tcp6
		router:     &Router{},
	}
}

//Server绑定一个方法
func (s *Server) Start() {
	fmt.Println("Server start...")

	addr := fmt.Sprintf("%s:%d", s.IP, s.Port)
	//创建socket，监听
	tcpaddr, err := net.ResolveTCPAddr(s.TCPVersion, addr) //调用函数，生成固定格式的addr，传入监听函数中
	if err != nil {
		fmt.Println("ResolveTCPAddr err", err)
		return
	}
	TCPListener, err := net.ListenTCP(s.TCPVersion, tcpaddr) //启动监听
	if err != nil {
		fmt.Println("ListenTCP err", err)
		return
	}

	var connId uint32
	//建立连接,Accept
	go func() {
		for {
			Tcpconn, err := TCPListener.AcceptTCP()
			if err != nil {
				fmt.Println("AcceptTCP err", err)
				return
			}
			fmt.Println("连接建立成功")

			//调用原生connection
			connId++
			conn := Newconn(Tcpconn, connId, s.router) //创建原生connection

			//对conn进行处理，接收client，转换成大写返回
			//对读写功能进行封装到Start函数中，用Connection调用即可
			go conn.Start()

		}
	}()

}

func (s *Server) Stop() {
	fmt.Println("server stop...")
}

func (s *Server) Server() {
	fmt.Println("server server...")
	s.Start()
	select {}
}

//添加路由方法
func (s *Server) AddRouter(router iface.IRouter) {
	//将传入的router，用Server中的字段接收
	s.router = router
}
