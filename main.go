package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"grpc-pixiu/service"
)

func main() {
	server := grpc.NewServer()

	service.RegisterProdServiceServer(server, service.ProductService)

	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("启动监听出错：", err)
	}
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("启动服务出错：", err)
	}
	fmt.Println("启动grpc服务端成功")
}
