package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"grpc-pixiu/service"
)

func WriteFile(request *service.ClusterRequest) error {
	clusterConfiguration := GetConfig()
	globalsFile := ClusterConfiguration{
		Kube_Release:            request.MasterInfo["HostName"],
		ClusterCidr:             request.MasterInfo["Username"],
		ServiceCidr:             request.MasterInfo["Password"],
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
	if err = ioutil.WriteFile(File, newfile, 0777); err != nil {
		log.Fatal("xml写入错误")
	}
	return err
}
