package main

import (
	"fmt"
	"grpc/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()

	service.RegisterProdServiceServer(rpcServer, service.ProductService)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("启动监听出错", err)
	}

	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("启动服务出错", err)
	}

	fmt.Println("启动 grpc 服务端成功")
}
