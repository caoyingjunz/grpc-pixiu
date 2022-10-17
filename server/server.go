package main

import (
	"net"

	"google.golang.org/grpc"
	"k8s.io/klog"

	"grpc-pixiu/cmd"
	pixiupb "grpc-pixiu/pixiu"
)

func main() {
	newServer := grpc.NewServer()

	pixiupb.RegisterCreateServiceServer(newServer, cmd.CreateClusterService)

	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		klog.Error("tcp listen fail:", err)
	}
	err = newServer.Serve(listen)
	if err != nil {
		klog.Error("start service fail:", err)
	}
}
