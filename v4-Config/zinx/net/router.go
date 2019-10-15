package net

import (
	"Zinx/Project/Zinx/v4-Config/zinx/iface"
	"fmt"
)

type Router struct {

}

func(r *Router) Handle(iface.IRequest)  {
	fmt.Println("Zinx Router-Handle ")
}
