package service

import (
	"context"
	"log"
	"os/exec"

	"grpc-pixiu/options"
)

var CreateClusterService = &createClusterService{}

type createClusterService struct {
}

func (p *createClusterService) Check(ctx context.Context, clusterInfo *ClusterRequest) (*ClusterResponse, error) {
	var err error
	startTime := options.GetStartTime()

	// 检查系统内是否安装kubez-ansible
	checkCmd := exec.Command(options.CheckKubezCommand)
	err = checkCmd.Run()
	if err != nil {
		installCmd := exec.Command("/bin/bash", "-c", options.InstallKubezCommand)
		err := installCmd.Run()
		if err == nil {
			multinodeCheckCmd := exec.Command(options.MultinodeCheckCmd)
			if err = multinodeCheckCmd.Run(); err != nil {
				log.Fatal("Multinode 配置:", err)
			}
		}
		log.Fatal("precondition fail:", err)
	}
	endTime := options.GetEndTime()
	return &ClusterResponse{
		ResponseInfo: "precondition successful:",
		StartTime:    startTime,
		EndTime:      endTime,
	}, err
}
