package client

import (
	"context"
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/kim-sardine/kfadmin/client/manifest"

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
func (c *KfClient) UpdateConfigMap(namespace, name string, dc manifest.DexConfigManifest) error {
	cm := c.GetConfigMap("auth", "dex")
	cm.Data["config.yaml"] = manifest.MarshalDexConfig(dc)
	_, err := c.cs.CoreV1().ConfigMaps(namespace).Update(context.TODO(), cm, metav1.UpdateOptions{})
	return err
}

// GetStaticUsers TBU
func (c *KfClient) GetStaticUsers() ([]manifest.StaticPasswordManifest, error) {
	cm := c.GetConfigMap("auth", "dex")
	data := cm.Data["config.yaml"]
	var dc manifest.DexConfigManifest
	err := yaml.Unmarshal([]byte(data), &dc)
	if err != nil {
		return []manifest.StaticPasswordManifest{}, err
	}
	return dc.StaticPasswords, nil
}

// GetProfile TBU
func (c *KfClient) GetProfile(profileName string) (manifest.Profile, error) {
	data, err := c.cs.RESTClient().
		Get().
		AbsPath("/apis/kubeflow.org/v1/profiles").
		Name(profileName).
		DoRaw(context.TODO())
	if err != nil {
		return manifest.Profile{}, err
	}

	profile, err := manifest.UnmarshalProfile(data)
	if err != nil {
		return manifest.Profile{}, err
	}

	return profile, nil
}

// CreateProfile TBU
func (c *KfClient) CreateProfile(profile manifest.Profile) error {
	body, err := manifest.MarshalProfile(profile)
	if err != nil {
		return err
	}

	_, err = c.cs.RESTClient().
		Post().
		AbsPath("/apis/kubeflow.org/v1/profiles").
		Body(body).
		DoRaw(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

// RestartDexDeployment TBU
// TODO: Should we restart dex automatically? or let admin manually restart it?
// https://www.kubeflow.org/docs/started/k8s/kfctl-istio-dex/#add-static-users-for-basic-auth
func (c *KfClient) RestartDexDeployment(backupData string) error {
	cmd := exec.Command("kubectl", "rollout", "restart", "deployment", "dex", "-n", "auth")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("failed to restart dex deployment")

		fmt.Println("start rollback dex configmap...")
		dc := manifest.UnmarshalDexConfig(backupData)
		err2 := c.UpdateConfigMap("auth", "dex", dc)
		if err2 != nil {
			return err2
		}
		fmt.Println("finish rollback dex configmap")
		return err
	}
	return nil
}
