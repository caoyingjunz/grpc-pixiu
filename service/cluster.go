package service

import (
	"context"
	"fmt"
	"grpc-pixiu/types"
	"os/exec"
)

var ProductService = &productService{}

type productService struct {
}

func CheckKubez() {
	checkKUbezcommand := exec.Command(types.CheckKubezCommand)
	err := checkKUbezcommand.Run()
	if err != nil {
		installKubezCommand := exec.Command("/bin/bash", "-c", types.InstallKubezCommand)
		fmt.Println("11111111")
		err := installKubezCommand.Run()
		if err != nil {
			fmt.Println("安装kubez出现问题", err)
		}
		fmt.Println("继续安装")
	}

}

func (p *productService) GetProductStock(context context.Context, request *ProductRequest) (*ProductResponse, error) {
	// 具体实现业务逻辑

	CheckKubez()
	stock, _ := p.GetStockById(request.ClusterName, request.MasterInfo)

	// 建立SSH客户端连接

	return &ProductResponse{Clustername: stock}, nil
}

func (p *productService) GetStockById(ClusterName string, username map[string]string) (string, map[string]string) {

	return ClusterName, username
}
