package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-pixiu/service"
)

func main() {
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("服务端出错，连接不上", err)
	}
	defer conn.Close()

	client := service.NewProdServiceClient(conn)
	request := &service.ProductRequest{
		ClusterName: "huawei",
		MasterInfo: map[string]string{
			"HostName": "81.68.210.233",
			"Adress":   "81.68.210.233",
			"Username": "root",
			"Password": "qq13069139214.",
		},
	}

	stockResponse, err := client.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal("查询库存出错", err)
	}
	fmt.Println("查询成功", stockResponse)
}
