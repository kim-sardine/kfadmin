package manifest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testNamespace = "test-ns"
	testUserName  = "tester@test.test"
)

func TestGetBindingName(t *testing.T) {
	bindingName, _ := GetBindingName(testUserName)

	assert.Equal(t, "user-tester-test-test-clusterrole-edit", bindingName)
}

func TestGetRoleBinding(t *testing.T) {
	rb, _ := GetRoleBinding(testNamespace, testUserName)
	bindingName, _ := GetBindingName(testUserName)

	assert.Equal(t, "rbac.authorization.k8s.io/v1", rb.TypeMeta.APIVersion)
	assert.Equal(t, "RoleBinding", rb.TypeMeta.Kind)
	assert.Equal(t, testUserName, rb.ObjectMeta.Annotations["user"])
	assert.Equal(t, testNamespace, rb.ObjectMeta.Namespace)
	assert.Equal(t, bindingName, rb.ObjectMeta.Name)
	assert.Equal(t, "edit", rb.ObjectMeta.Annotations["role"])
	assert.Equal(t, "ClusterRole", rb.RoleRef.Kind)
	assert.Equal(t, "kubeflow-edit", rb.RoleRef.Name)

}

func TestGetServiceRoleBinding(t *testing.T) {
	rb, _ := GetServiceRoleBinding(testNamespace, testUserName)
	bindingName, _ := GetBindingName(testUserName)

	assert.Equal(t, "rbac.istio.io/v1alpha1", rb.TypeMeta.APIVersion)
	assert.Equal(t, "ServiceRoleBinding", rb.TypeMeta.Kind)
	assert.Equal(t, testUserName, rb.ObjectMeta.Annotations["user"])
	assert.Equal(t, testNamespace, rb.ObjectMeta.Namespace)
	assert.Equal(t, bindingName, rb.ObjectMeta.Name)
	assert.Equal(t, "edit", rb.ObjectMeta.Annotations["role"])
	assert.Equal(t, "ServiceRole", rb.Spec.RoleRef.Kind)
	assert.Equal(t, "ns-access-istio", rb.Spec.RoleRef.Name)
	assert.Equal(t, testUserName, rb.Spec.Subjects[0].Properties["request.headers[kubeflow-userid]"])
}
