package iface

type IRequest interface {
	GetLen() uint32
	GetData() []byte
	GetConn() Iconnection
}
