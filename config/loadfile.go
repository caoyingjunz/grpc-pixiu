package config

import (
	"os"

	"gopkg.in/yaml.v2"
	"k8s.io/klog"
)

const YmlFile = "/etc/kubez/globals.yml"

func loadConfigFromFile(file string) (*BuildCloud, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return loadConfig(data)
}

func loadConfig(data []byte) (*BuildCloud, error) {
	var pc BuildCloud
	if err := yaml.Unmarshal(data, &pc); err != nil {
		return nil, err
	}

	return &pc, nil
}

func GetConfig() *BuildCloud {
	clusterConfiguration, err := loadConfigFromFile(YmlFile)
	if err != nil {
		klog.Error("load config fail:", err)
	}
	return clusterConfiguration
}
