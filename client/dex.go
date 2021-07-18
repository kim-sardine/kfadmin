package client

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/kim-sardine/kfadmin/manifest"

	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetConfigMap TBU
func (c *KfClient) GetConfigMap(namespace, name string) (*v1.ConfigMap, error) {
	return c.cs.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

// GetDexConfigMap TBU
func (c *KfClient) GetDexConfigMap() (*v1.ConfigMap, error) {
	dex, err := c.GetConfigMap("auth", "dex")
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, fmt.Errorf("can't find configmap \"dex\"\nThis command only works for dex")
		}
		return nil, err
	}
	return dex, nil
}

// UpdateConfigMap TBU
func (c *KfClient) UpdateConfigMap(namespace string, cm *v1.ConfigMap) error {
	_, err := c.cs.CoreV1().ConfigMaps(namespace).Update(context.TODO(), cm, metav1.UpdateOptions{})
	return err
}

// UpdateDexConfigMap TBU
func (c *KfClient) UpdateDexConfigMap(cm *v1.ConfigMap) error {
	err := c.UpdateConfigMap("auth", cm)
	return err
}

// GetStaticUsers TBU
func (c *KfClient) GetStaticUsers() ([]manifest.StaticPassword, error) {
	cm, err := c.GetDexConfigMap()
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

// RollbackDexDeployment TBU
func (c *KfClient) RollbackDexDeployment(backupData string) error {
	cm, err := c.GetDexConfigMap()
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

	if err2 := c.UpdateDexConfigMap(cm); err2 != nil {
		return err2
	}

	return nil
}

// RestartDexDeployment TBU
// https://www.kubeflow.org/docs/started/k8s/kfctl-istio-dex/#add-static-users-for-basic-auth
func (c *KfClient) RestartDexDeployment() error {
	dexDeployment, err := c.cs.AppsV1().Deployments("auth").Get(context.TODO(), "dex", metav1.GetOptions{})
	if err != nil {
		return err
	}

	dexDeployment.Spec.Template.Annotations["date"] = time.Now().String()
	_, err = c.cs.AppsV1().Deployments("auth").Update(context.TODO(), dexDeployment, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (c *KfClient) RestartDex(out io.Writer, backupData string) error {
	if err := c.RestartDexDeployment(); err != nil {
		fmt.Fprintf(out, err.Error()+"\n")
		fmt.Fprintf(out, "failed to restart dex deployment, rollback dex deployment...\n")
		if err = c.RollbackDexDeployment(backupData); err != nil {
			return err
		}
		return fmt.Errorf("completed rollback dex deployment")
	}
	return nil
}
