package client

import (
	"context"
	"encoding/json"

	"github.com/kim-sardine/kfadmin/manifest"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
