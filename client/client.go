package client

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/kim-sardine/kfadmin/client/manifest"

	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	GetProfileList() (manifest.ProfileList, error)
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

// Kubeflow Profile
// https://www.kubeflow.org/docs/components/multi-tenancy/getting-started/#manual-profile-creation

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

// GetProfileList TBU
func (c *KfClient) GetProfileList() (manifest.ProfileList, error) {
	data, err := c.cs.RESTClient().
		Get().
		AbsPath("/apis/kubeflow.org/v1/profiles").
		DoRaw(context.TODO())
	if err != nil {
		return manifest.ProfileList{}, err
	}

	profileList, err := manifest.UnmarshalProfileList(data)
	if err != nil {
		return manifest.ProfileList{}, err
	}

	return profileList, nil
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

// UpdateProfile TBU
func (c *KfClient) UpdateProfile(profile manifest.Profile) error {
	body, err := manifest.MarshalProfile(profile)
	if err != nil {
		return err
	}

	_, err = c.cs.RESTClient().
		Put().
		AbsPath("/apis/kubeflow.org/v1/profiles").
		Name(profile.ObjectMeta.Name).
		Body(body).
		DoRaw(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

// DeleteProfile TBU
func (c *KfClient) DeleteProfile(profileName string) error {
	_, err := c.cs.RESTClient().
		Delete().
		AbsPath("/apis/kubeflow.org/v1/profiles").
		Name(profileName).
		DoRaw(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

// CreateRoleBinding TBU
func (c *KfClient) CreateRoleBinding(namespace string, roleBinding *rbacv1.RoleBinding) error {
	_, err := c.cs.RbacV1().RoleBindings(namespace).Create(context.TODO(), roleBinding, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

// GetRoleBinding TBU
func (c *KfClient) GetRoleBinding(namespace, name string) (*rbacv1.RoleBinding, error) {
	rb, err := c.cs.RbacV1().RoleBindings(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return rb, nil
}

// UpdateRoleBinding TBU
func (c *KfClient) UpdateRoleBinding(namespace string, roleBinding *rbacv1.RoleBinding) error {
	_, err := c.cs.RbacV1().RoleBindings(namespace).Update(context.TODO(), roleBinding, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoleBinding TBU
func (c *KfClient) DeleteRoleBinding(namespace string, name string) error {
	err := c.cs.RbacV1().RoleBindings(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	return err
}

// GetServiceRoleBinding TBU
func (c *KfClient) GetServiceRoleBinding(namespace string, name string) (*manifest.ServiceRoleBinding, error) {
	absPath := "/apis/rbac.istio.io/v1alpha1/namespaces/" + namespace + "/servicerolebindings"

	data, err := c.cs.RESTClient().
		Get().
		AbsPath(absPath).
		Name(name).
		DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}

	var srb manifest.ServiceRoleBinding
	err = json.Unmarshal(data, &srb)
	if err != nil {
		return nil, err
	}
	return &srb, nil
}

// CreateServiceRoleBinding TBU
func (c *KfClient) CreateServiceRoleBinding(namespace string, serviceRoleBinding *manifest.ServiceRoleBinding) error {
	data, err := json.Marshal(serviceRoleBinding)
	if err != nil {
		return err
	}

	absPath := "/apis/rbac.istio.io/v1alpha1/namespaces/" + namespace + "/servicerolebindings"

	_, err = c.cs.RESTClient().
		Post().
		AbsPath(absPath).
		Body(data).
		DoRaw(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

// UpdateServiceRoleBinding TBU
func (c *KfClient) UpdateServiceRoleBinding(namespace string, serviceRoleBinding *manifest.ServiceRoleBinding) error {
	data, err := json.Marshal(serviceRoleBinding)
	if err != nil {
		return err
	}

	absPath := "/apis/rbac.istio.io/v1alpha1/namespaces/" + namespace + "/servicerolebindings"

	_, err = c.cs.RESTClient().
		Put().
		AbsPath(absPath).
		Name("owner-binding-istio").
		Body(data).
		DoRaw(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

// DeleteServiceRoleBinding TBU
func (c *KfClient) DeleteServiceRoleBinding(namespace string, name string) error {
	absPath := "/apis/rbac.istio.io/v1alpha1/namespaces/" + namespace + "/servicerolebindings"

	_, err := c.cs.RESTClient().
		Delete().
		AbsPath(absPath).
		Name(name).
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
