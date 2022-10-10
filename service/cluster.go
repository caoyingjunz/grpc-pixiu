package service

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

var ProductService = &productService{}

type productService struct {
}

func (p *productService) GetProductStock(context context.Context, request *ProductRequest) (*ProductResponse, error) {
	// 具体实现业务逻辑
	stock, name := p.GetStockById(request.ClusterName, request.MasterInfo)
	fmt.Println(stock, name, request.MasterInfo)

	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", request.MasterInfo["Adress"]+":22", &ssh.ClientConfig{
		User:            request.MasterInfo["Username"],
		Auth:            []ssh.AuthMethod{ssh.Password(request.MasterInfo["Password"])},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatal("建立客户端连接错误", err)
	}

	// 建立新会话
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("建立 session 错误", err)
	}
	//result, err := session.Output("curl -# -O https://raw.githubusercontent.com/caoyingjunz/kubez-ansible/master/tools/setup_env.sh " +
	//	"&& date " +
	//	"&& bash ~/setup_env.sh " +
	//	"&& kubez-ansible bootstrap-servers " +
	//	"&& kubez-ansible deploy " +
	//	"&& kubez-ansible post-deploy " +
	//	"&& kubectl get node")
	_, err = session.Output("ls -al && date")

	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to run command, Err:%s", err.Error())
		os.Exit(0)
	}

	return &ProductResponse{Clustername: stock}, nil
}

func (p *productService) GetStockById(ClusterName string, username map[string]string) (string, map[string]string) {

	return ClusterName, username
}
