package iface

type IRouter interface {
	Handle(IRequest)
}
