package cmd

import (
	"context"
	"os/exec"

	"k8s.io/klog"

	"grpc-pixiu/options"
	pixiupb "grpc-pixiu/pixiu"
)

func (p *createClusterService) Installation(ctx context.Context, clusterInfo *pixiupb.ClusterRequest) (*pixiupb.ClusterResponse, error) {
	var err error
	clusterName := clusterInfo.Name
	startTime := options.GetStartTime()

	multinodeInstallCmd := exec.Command("/bin/bash", "-c", options.MultinodeInstallCmd)
	if err = multinodeInstallCmd.Run(); err != nil {
		klog.Errorf("multinode install:", err)
	}

	deployCmd := exec.Command("/bin/bash", "-c", options.DeployCmd)
	if err = deployCmd.Run(); err != nil {
		klog.Errorf("deploy configuration:", err)
	}

	rcCmd := exec.Command("/bin/bash", "-c", options.RcCmd)
	if err = rcCmd.Run(); err != nil {
		klog.Errorf("rc configuration:", err)
	}
	endTime := options.GetEndTime()

	return &pixiupb.ClusterResponse{
		ResponseInfo: "Kubernetes installation successful:" + clusterName,
		StartTime:    startTime,
		EndTime:      endTime,
	}, nil
}
