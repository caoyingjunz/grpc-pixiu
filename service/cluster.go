package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"gopkg.in/yaml.v2"

	"grpc-pixiu/config"
	"grpc-pixiu/types"
)

var CreateClusterService = &createClusterService{}

type createClusterService struct {
}

func CheckKubez(request *ClusterRequest) error {
	var error error

	// 检查系统内是否安装kubez-ansible
	checkCmd := exec.Command(types.CheckKubezCommand)
	error = checkCmd.Run()
	if error != nil {
		installCmd := exec.Command("/bin/bash", "-c", types.InstallKubezCommand)
		error := installCmd.Run()
		if error == nil {
			multinodeCmd := exec.Command("ansible -i multinode all -m ping")
			if error = multinodeCmd.Run(); error != nil {
				fmt.Println("Multinode 配置出错", error)
			}

			BindXml(request)

			deployCmd := exec.Command("kubez-ansible -i multinode deploy")
			if error = deployCmd.Run(); error != nil {
				fmt.Println("deploy 配置出错", error)
			}

			rcCmd := exec.Command("kubez-ansible -i multinode post-deploy")
			if error = rcCmd.Run(); error != nil {
				fmt.Println("rc 配置出错", error)
			}

			//fmt.Println("集群安装成功")
		}

		fmt.Println("安装kubez出现问题", error)

	}
	return error
}

func BindXml(request *ClusterRequest) {
	clusterConfiguration := config.GetConfig()
	globalsFile := config.ClusterConfiguration{
		Kube_Release:            request.MasterInfo["HostName"],
		ClusterCidr:             request.MasterInfo["Username"],
		ServiceCidr:             request.MasterInfo["Password"],
		DockerRelease:           clusterConfiguration.DockerRelease,
		DockerReleaseUbuntu:     clusterConfiguration.DockerReleaseUbuntu,
		ContainerdRelease:       clusterConfiguration.ContainerdRelease,
		ContainerdReleaseUbuntu: clusterConfiguration.ContainerdReleaseUbuntu,
	}
	out, err := yaml.Marshal(globalsFile)
	if err != nil {
		log.Fatal("xml解析错误")
	}
	// 通过ioutil.writeFile写入文件
	err = ioutil.WriteFile(config.File, out, 0777)
	if err != nil {
		log.Fatal("xml写入错误")
	}
}

func (p *createClusterService) CreateCluster(context context.Context, request *ClusterRequest) (*ClusterResponse, error) {
	clusterName := request.ClusterName

	err := CheckKubez(request)
	if err != nil {
		return &ClusterResponse{ClusterName: clusterName}, err
	}

	return &ClusterResponse{ClusterName: clusterName}, nil
}
