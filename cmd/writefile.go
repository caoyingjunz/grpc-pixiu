package cmd

import (
	"context"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"k8s.io/klog"

	"grpc-pixiu/config"
	"grpc-pixiu/options"
	pixiupb "grpc-pixiu/pixiu"
)

func (p *createClusterService) WriteFile(ctx context.Context, clusterInfo *pixiupb.ClusterRequest) (*pixiupb.ClusterResponse, error) {
	startTime := options.GetStartTime()

	clusterConfiguration := config.GetConfig()
	globalsFile := config.BuildCloud{
		Name:      clusterInfo.Name,
		AliasName: clusterConfiguration.AliasName,
		// TODO继续增加渲染的参数
	}
	newfile, err := yaml.Marshal(globalsFile)
	if err != nil {
		klog.Errorf("parsing xml fail:,", err)
	}
	// 通过ioutil.writeFile写入文件
	if err = ioutil.WriteFile(config.File, newfile, 0777); err != nil {
		klog.Errorf("write xml fail:,", err)
	}
	endTime := options.GetEndTime()
	return &pixiupb.ClusterResponse{
		ResponseInfo: "write globals file successful:",
		StartTime:    startTime,
		EndTime:      endTime,
	}, err
}
