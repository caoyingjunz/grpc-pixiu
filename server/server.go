package main

import (
	"grpc-pixiu/cmd"
	"log"
	"net"

	"google.golang.org/grpc"

	pixiupb "grpc-pixiu/pixiu"
)

func main() {
	newServer := grpc.NewServer()

	pixiupb.RegisterCreateServiceServer(newServer, cmd.CreateClusterService)

	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("启动监听出错：", err)
	}
	err = newServer.Serve(listen)
	if err != nil {
		log.Fatal("启动服务出错：", err)
	}
}
