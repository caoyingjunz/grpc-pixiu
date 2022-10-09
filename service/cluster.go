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
	stock, name, password, ip := p.GetStockById(request.ClusterName, request.UserName, request.Password, request.Ip)
	fmt.Println(stock, name, password)

	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", ip+":22", &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatal("建立客户端连接错误", err)
	}
	// 建立新会话
	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		log.Fatal("建立 session 错误", err)
	}
	result, err := session.Output("curl https://raw.githubusercontent.com/caoyingjunz/kubez-ansible/master/tools/setup_env.sh")

	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to run command, Err:%s", err.Error())
		os.Exit(0)
	}

	fmt.Println(string(result))

	return &ProductResponse{Clustername: stock, Username: name, Password: password}, nil

}

func (p *productService) GetStockById(ClusterName string, username string, password string, ip string) (string, string, string, string) {

	return ClusterName, username, password, ip
}
