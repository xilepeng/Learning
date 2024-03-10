package main

import (
	"context"
	"fmt"
	"grpc/client/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. 新建连接
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("服务端出错，连接不上", err)
	}

	// 2. 退出关闭连接
	defer conn.Close()

	prodClient := service.NewProdServiceClient(conn)

	request := &service.ProductRequest{ProdId: 123}
	stockResponse, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal("查询库存出错", err)
	}

	fmt.Println("查询成功", stockResponse.ProdStock)

}
