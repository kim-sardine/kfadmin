package client

import (
	"flag"
	"path/filepath"

	"github.com/kim-sardine/kfadmin/manifest"

	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Client TBU
type Client interface {
	LoadClientset()
	GetConfigMap(namespace, name string) (*v1.ConfigMap, error)
	GetDex() (*v1.ConfigMap, error)
	UpdateConfigMap(namespace string, cm *v1.ConfigMap) error
	UpdateDex(cm *v1.ConfigMap) error
	GetStaticUsers() ([]manifest.StaticPassword, error)
	GetProfile(profileName string) (manifest.Profile, error)
	GetProfiles() (manifest.Profiles, error)
	CreateProfile(profile manifest.Profile) error
	UpdateProfile(profile manifest.Profile) error
	DeleteProfile(profileName string) error
	CreateRoleBinding(namespace string, roleBinding *rbacv1.RoleBinding) error
	GetRoleBinding(namespace, name string) (*rbacv1.RoleBinding, error)
	UpdateRoleBinding(namespace string, roleBinding *rbacv1.RoleBinding) error
	DeleteRoleBinding(namespace string, name string) error
	GetServiceRoleBinding(namespace string, name string) (*manifest.ServiceRoleBinding, error)
	CreateServiceRoleBinding(namespace string, serviceRoleBinding *manifest.ServiceRoleBinding) error
	UpdateServiceRoleBinding(namespace string, serviceRoleBinding *manifest.ServiceRoleBinding) error
	DeleteServiceRoleBinding(namespace string, name string) error
	RestartDexDeployment(backupData string) error
}

// KfClient TBU
type KfClient struct {
	cs *kubernetes.Clientset
}

func NewKfClient() *KfClient {
	return &KfClient{}
}

// LoadClientset TBU
func (c *KfClient) LoadClientset() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	c.cs = clientset
}
