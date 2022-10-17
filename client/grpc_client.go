package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pixiupb "grpc-pixiu/pixiu"
)

const address = "localhost:8002"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("服务端出错，连接不上", err)
	}
	defer conn.Close()

	client := pixiupb.NewCreateServiceClient(conn)
	createCluster := &pixiupb.ClusterRequest{
		Name:      "huawei",
		AliasName: "test",
	}

	//check, err := client.Check(context.Background(), createCluster)
	//if err != nil {
	//	log.Fatal("precondition fail", err)
	//}
	//fmt.Println("precondition successful", check)

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
