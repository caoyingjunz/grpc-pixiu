package cmd

import (
	"context"
	"log"
	"os/exec"

	"grpc-pixiu/options"
	pixiupb "grpc-pixiu/pixiu"
)

func (p *createClusterService) Installation(ctx context.Context, clusterInfo *pixiupb.ClusterRequest) (*pixiupb.ClusterResponse, error) {
	var err error
	clusterName := clusterInfo.Name
	startTime := options.GetStartTime()

	multinodeInstallCmd := exec.Command("/bin/bash", "-c", options.MultinodeInstallCmd)
	if err = multinodeInstallCmd.Run(); err != nil {
		log.Fatal("Multinode install 配置:", err)
	}

	deployCmd := exec.Command("/bin/bash", "-c", options.DeployCmd)
	if err = deployCmd.Run(); err != nil {
		log.Fatal("deploy 配置:", err)
	}

	rcCmd := exec.Command("/bin/bash", "-c", options.RcCmd)
	if err = rcCmd.Run(); err != nil {
		log.Fatal("rc 配置:", err)
	}
	endTime := options.GetEndTime()

	return &pixiupb.ClusterResponse{
		ResponseInfo: "Kubernetes installation successful:" + clusterName,
		StartTime:    startTime,
		EndTime:      endTime,
	}, nil
}
