package options

const (
	InstallKubezCommand = "curl https://raw.githubusercontent.com/caoyingjunz/kubez-ansible/master/tools/setup_env.sh"
	CheckKubezCommand   = "kubez-ansible"
	MultinodeCheckCmd   = "ansible -i multinode all -m ping"
	MultinodeInstallCmd = "kubez-ansible bootstrap-servers"
	DeployCmd           = "kubez-ansible deploy"
	RcCmd               = "kubez-ansible post-deploy"
)
