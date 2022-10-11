package config

type ClusterConfiguration struct {
	Kube_Release            string `yaml:"kube_release"`
	ClusterCidr             string `yaml:"cluster_cidr"`
	ServiceCidr             string `yaml:"service_cidr"`
	DockerRelease           string `yaml:"docker_release"`
	DockerReleaseUbuntu     string `yaml:"docker_release_ubuntu"`
	ContainerdRelease       string `yaml:"containerd_release"`
	ContainerdReleaseUbuntu string `yaml:"containerd_release_ubuntu"`
}

type Galbals map[string]interface{}
