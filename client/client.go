package client

import (
	"context"
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Client TBU
type Client interface {
	LoadClientset()
	ListPods(string)
	GetConfigMap(string, string) *v1.ConfigMap
}

// KfClient TBU
type KfClient struct {
	cs *kubernetes.Clientset
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

// ListPods TBU
func (c *KfClient) ListPods(namespace string) {
	pods, err := c.cs.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(pods)
}

// GetConfigMap TBU
func (c *KfClient) GetConfigMap(namespace, name string) *v1.ConfigMap {
	cm, err := c.cs.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	return cm
}

// UpdateConfigMap TBU
func (c *KfClient) UpdateConfigMap(namespace, name string, cm *v1.ConfigMap) error {
	_, err := c.cs.CoreV1().ConfigMaps(namespace).Update(context.TODO(), cm, metav1.UpdateOptions{})
	return err
}

// GetStaticUsers TBU
func (c *KfClient) GetStaticUsers() []StaticPasswordManifest {
	cm := c.GetConfigMap("auth", "dex")
	data := cm.Data["config.yaml"]
	var dc DexConfigManifest
	err := yaml.Unmarshal([]byte(data), &dc)
	if err != nil {
		panic(err)
	}
	return dc.StaticPasswords
}

// RestartDexDeployment TBU
func (c *KfClient) RestartDexDeployment() error {
	cmd := exec.Command("kubectl", "rollout", "restart", "deployment", "dex", "-n", "auth")
	_, err := cmd.Output()
	return err
}
