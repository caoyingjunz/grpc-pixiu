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
		testMultinode := exec.Command("ansible -i multinode all -m ping")
		err = testMultinode.Run()
		if err != nil {
			fmt.Println("Multinode 配置出错", err)
		}
		// TODO 渲染xml文件
		// TODO 处理XML文件
		deploy := exec.Command("kubez-ansible -i multinode deploy")
		err = deploy.Run()
		if err != nil {
			fmt.Println("deploy 配置出错", err)
		}
		rc := exec.Command("kubez-ansible -i multinode post-deploy")
		err = rc.Run()
		if err != nil {
			fmt.Println("rc 配置出错", err)
		}
		fmt.Println("集群安装成功")
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
