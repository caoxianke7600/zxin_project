package iface

//创建Server接口
type Iserver interface {
	Start()
	Stop()
	Server()
	AddRouter(router IRouter)   //增加单个路由的方法
}
