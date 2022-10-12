package options

const (
	InstallKubezCommand = "curl https://raw.githubusercontent.com/caoyingjunz/kubez-ansible/master/tools/setup_env.sh"
	CheckKubezCommand   = "kubez-ansible"
	MultinodeCheckCmd   = "ansible -i multinode all -m ping"
	MultinodeInstallCmd = "kubez-ansible -i multinode bootstrap-servers"
	DeployCmd           = "kubez-ansible -i multinode deploy"
	RcCmd               = "kubez-ansible -i multinode post-deploy"
)
