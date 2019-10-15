package main

import "Zinx/Project/Zinx/v3-Request/zinx/net"

func main() {
	server := net.NewServer("Zinxv1.0")
	server.Server()
}
