package client

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/kim-sardine/kfadmin/manifest"

	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetConfigMap TBU
func (c *KfClient) GetConfigMap(namespace, name string) (*v1.ConfigMap, error) {
	cm, err := c.cs.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	return cm, err
}

// GetDex TBU
func (c *KfClient) GetDex() (*v1.ConfigMap, error) {
	dex, err := c.GetConfigMap("auth", "dex")
	if err != nil {
		return nil, err
	}
	return dex, nil
}

// UpdateConfigMap TBU
func (c *KfClient) UpdateConfigMap(namespace string, cm *v1.ConfigMap) error {
	_, err := c.cs.CoreV1().ConfigMaps(namespace).Update(context.TODO(), cm, metav1.UpdateOptions{})
	return err
}

// UpdateDex TBU
func (c *KfClient) UpdateDex(cm *v1.ConfigMap) error {
	err := c.UpdateConfigMap("auth", cm)
	return err
}

// GetStaticUsers TBU
func (c *KfClient) GetStaticUsers() ([]manifest.StaticPassword, error) {
	cm, err := c.GetDex()
	if err != nil {
		return []manifest.StaticPassword{}, err
	}

	data := cm.Data["config.yaml"]
	var dc manifest.DexDataConfig
	err = yaml.Unmarshal([]byte(data), &dc)
	if err != nil {
		return []manifest.StaticPassword{}, err
	}
	return dc.StaticPasswords, nil
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
		cm, err := c.GetDex()
		if err != nil {
			return err
		}

		dc, err := manifest.UnmarshalDexDataConfig(backupData)
		if err != nil {
			return err
		}
		cm.Data["config.yaml"], err = manifest.MarshalDexDataConfig(dc)
		if err != nil {
			return err
		}

		err2 := c.UpdateDex(cm)
		if err2 != nil {
			return err2
		}
		fmt.Println("finish rollback dex configmap")
		return err
	}
	return nil
}
