package cmd

import (
	"log"
	"os/exec"

	"grpc-pixiu/config"
	"grpc-pixiu/service"
	"grpc-pixiu/types"
)

func CheckKubez(request *service.ClusterRequest) error {
	var err error

	// 检查系统内是否安装kubez-ansible
	checkCmd := exec.Command(types.CheckKubezCommand)
	err = checkCmd.Run()
	if err != nil {
		installCmd := exec.Command("/bin/bash", "-c", types.InstallKubezCommand)
		err := installCmd.Run()
		if err == nil {
			multinodeCheckCmd := exec.Command(types.MultinodeCheckCmd)
			if err = multinodeCheckCmd.Run(); err != nil {
				log.Fatal("Multinode 配置:", err)
			}

			// 处理客户端传来的数据，重新写入到globals.yml
			if err := config.WriteFile(request); nil != nil {
				log.Fatal("deploy 配置:", err)
			}

			multinodeInstallCmd := exec.Command(types.MultinodeInstallCmd)
			if err = multinodeInstallCmd.Run(); err != nil {
				log.Fatal("Multinode install 配置:", err)
			}

			deployCmd := exec.Command(types.DeployCmd)
			if err = deployCmd.Run(); err != nil {
				log.Fatal("deploy 配置:", err)
			}

			rcCmd := exec.Command(types.RcCmd)
			if err = rcCmd.Run(); err != nil {
				log.Fatal("rc 配置:", err)
			}

		}

		log.Fatal("安装kubez出现问题", err)

	}
	return err
}
