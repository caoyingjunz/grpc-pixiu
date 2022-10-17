package config

type BuildCloud struct {
	Name            string          `json:"name"`       // 名称，系统自动生成，只能为字符串
	AliasName       string          `json:"alias_name"` // 可读性的名称，支持中午
	Immediate       bool            `json:"immediate"`  // 立刻部署
	CloudType       int             `json:"cloud_type"` // cloud 的类型，支持标准类型和自建类型
	Region          string          `json:"region"`     // 城市区域
	Kubernetes      *KubernetesSpec `json:"kubernetes"` // k8s 全部信息
	CreateNamespace bool            `json:"create_namespace"`
	Description     string          `json:"description"`
}

type KubernetesSpec struct {
	ApiServer   string     `json:"api_server"` // kubernetes 的 apiServer 的 ip 地址
	Version     string     `json:"version"`    // k8s 的版本
	Runtime     string     `json:"runtime"`    // 容器运行时，目前支持 docker 和 containerd
	Cni         string     `json:"cni"`        // 网络 cni，支持 flannel 和 calico
	ServiceCidr string     `json:"service_cidr"`
	PodCidr     string     `json:"pod_cidr"`
	ProxyMode   string     `json:"proxy_mode"` // kubeProxy 的模式，只能是 iptables 和 ipvs
	Masters     []NodeSpec `json:"masters"`    // 集群的 master 节点
	Nodes       []NodeSpec `json:"nodes"`      // 集群的 node 节点
}

// Node k8s node属性
type Node struct {
	Name             string `json:"name"`
	Status           string `json:"status"`
	Roles            string `json:"roles"`
	CreateAt         string `json:"create_at"`
	Version          string `json:"version"`
	InternalIP       string `json:"internal_ip"`
	OsImage          string `json:"osImage"`
	KernelVersion    string `json:"kernel_version"`
	ContainerRuntime string `json:"container_runtime"`
}

// NodeSpec 构造 kubernetes 集群的节点
type NodeSpec struct {
	HostName string `json:"host_name"`
	Address  string `json:"address"`
	User     string `json:"user"`
	Password string `json:"password"`
}
