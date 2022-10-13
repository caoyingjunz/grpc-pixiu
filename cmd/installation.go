package cmd

import (
	"context"
	"log"
	"os/exec"

	"grpc-pixiu/options"
)

func (p *createClusterService) Installation(ctx context.Context, clusterInfo *ClusterRequest) (*ClusterResponse, error) {
	var err error
	clusterName := clusterInfo.ClusterName
	startTime := options.GetStartTime()

	multinodeInstallCmd := exec.Command(options.MultinodeInstallCmd)
	if err = multinodeInstallCmd.Run(); err != nil {
		log.Fatal("Multinode install 配置:", err)
	}

	deployCmd := exec.Command(options.DeployCmd)
	if err = deployCmd.Run(); err != nil {
		log.Fatal("deploy 配置:", err)
	}

	rcCmd := exec.Command(options.RcCmd)
	if err = rcCmd.Run(); err != nil {
		log.Fatal("rc 配置:", err)
	}
	endTime := options.GetEndTime()

	return &ClusterResponse{
		ResponseInfo: "Kubernetes installation successful:" + clusterName,
		StartTime:    startTime,
		EndTime:      endTime,
	}, nil
}
