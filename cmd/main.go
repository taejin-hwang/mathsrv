package main

import (
	"github.com/taejin-hwang/mathsrv/server"
)

func main() {
	server := server.NewServer("localhost", 8080)
	server.Start()
}