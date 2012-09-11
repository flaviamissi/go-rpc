package main

import (
	"github.com/flaviamissi/go-rpc/server"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	s := new(server.IaaS)
	rpc.Register(s)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		panic(e)
	}
	e = http.Serve(l, nil)
	if e != nil {
		panic(e)
	}
}
