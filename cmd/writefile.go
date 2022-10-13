package cmd

import (
	"context"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"grpc-pixiu/config"
	"grpc-pixiu/options"
)

func (p *createClusterService) WriteFile(ctx context.Context, clusterInfo *ClusterRequest) (*ClusterResponse, error) {
	startTime := options.GetStartTime()

	clusterConfiguration := config.GetConfig()
	globalsFile := config.ClusterConfiguration{
		Kube_Release:            clusterInfo.MasterInfo["HostName"],
		ClusterCidr:             clusterInfo.MasterInfo["Username"],
		ServiceCidr:             clusterInfo.MasterInfo["Password"],
		DockerRelease:           clusterConfiguration.DockerRelease,
		DockerReleaseUbuntu:     clusterConfiguration.DockerReleaseUbuntu,
		ContainerdRelease:       clusterConfiguration.ContainerdRelease,
		ContainerdReleaseUbuntu: clusterConfiguration.ContainerdReleaseUbuntu,
	}
	newfile, err := yaml.Marshal(globalsFile)
	if err != nil {
		log.Fatal("xml解析错误")
	}
	// 通过ioutil.writeFile写入文件
	if err = ioutil.WriteFile(config.File, newfile, 0777); err != nil {
		log.Fatal("xml写入错误")
	}
	endTime := options.GetEndTime()
	return &ClusterResponse{
		ResponseInfo: "write globals file successful:",
		StartTime:    startTime,
		EndTime:      endTime,
	}, err
}
