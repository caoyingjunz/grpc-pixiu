package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"grpc-pixiu/service"
)

const address = "localhost:8002"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("服务端出错，连接不上", err)
	}
	defer conn.Close()

	client := service.NewCreateServiceClient(conn)
	createCluster := &service.ClusterRequest{
		ClusterName: "huawei",
		MasterInfo: map[string]string{
			"HostName": "81.68.210.233",
			"Adress":   "81.68.210.233",
			"Username": "root",
			"Password": "qq13069139214.",
		},
	}

	check, err := client.Check(context.Background(), createCluster)
	if err != nil {
		log.Fatal("precondition fail", err)
	}
	fmt.Println("precondition successful", check)

	writefile, err := client.WriteFile(context.Background(), createCluster)
	if err != nil {
		log.Fatal("write globals file fail:", err)
	}
	fmt.Println("write globals file successful:", writefile)

	installation, err := client.Installation(context.Background(), createCluster)
	if err != nil {
		log.Fatal("Kubernetes installation fail", err)
	}
	fmt.Println("Kubernetes installation successful:", installation)
}
