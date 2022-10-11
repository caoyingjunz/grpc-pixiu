package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

const File = "globals.yml"

func loadConfigFromFile(file string) (*ClusterConfiguration, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return loadConfig(data)
}

func loadConfig(data []byte) (*ClusterConfiguration, error) {
	var pc ClusterConfiguration
	if err := yaml.Unmarshal(data, &pc); err != nil {
		return nil, err
	}

	return &pc, nil
}

func GetConfig() *ClusterConfiguration {
	clusterConfiguration, err := loadConfigFromFile(File)
	if err != nil {
		panic(err)
	}
	return clusterConfiguration
}
