package net

import "Zinx/Project/Zinx/v3-Request/zinx/iface"

type Request struct {
	conn iface.Iconnection
	data []byte

	len uint32
}

func NewRequest(conn iface.Iconnection, data []byte, len uint32) iface.IRequest {
	return &Request{
		conn: conn,
		data: data,
		len:  len,
	}
}

//实现方法
func (r *Request) GetConn() iface.Iconnection {
	return r.conn
}
func (r *Request) GetData() []byte {
	return r.data
}
func (r *Request) GetLen() uint32 {
	return r.len
}
