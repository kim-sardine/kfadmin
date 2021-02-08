package manifest

import (
	"net/url"
	"regexp"
	"strings"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetRoleBinding TBU
func GetRoleBinding(namespace, userName string) (*rbacv1.RoleBinding, error) {
	bindingName, err := GetBindingName(userName)
	if err != nil {
		return nil, err
	}
	rolebinding := &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{"user": userName, "role": "edit"},
			Name:        bindingName,
			Namespace:   namespace,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "kubeflow-edit",
		},
		Subjects: []rbacv1.Subject{
			{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "User",
				Name:     userName,
			},
		},
	}
	return rolebinding, nil
}

// GetServiceRoleBinding TBU
func GetServiceRoleBinding(namespace, userName string) (*ServiceRoleBinding, error) {
	bindingName, err := GetBindingName(userName)
	if err != nil {
		return nil, err
	}
	serviceRolebinding := &ServiceRoleBinding{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.istio.io/v1alpha1",
			Kind:       "ServiceRoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{"user": userName, "role": "edit"},
			Name:        bindingName,
			Namespace:   namespace,
		},
		Spec: ServiceRoleBindingSpec{
			RoleRef: &RoleRef{
				Kind: "ServiceRole",
				Name: "ns-access-istio",
			},
			Subjects: []*Subject{
				{
					Properties: map[string]string{"request.headers[kubeflow-userid]": userName},
				},
			},
		},
	}
	return serviceRolebinding, nil
}

// ServiceRoleBindingSpec defines the desired state of ServiceRoleBinding
type ServiceRoleBindingSpec struct {
	Subjects []*Subject `json:"subjects,omitempty"`
	RoleRef  *RoleRef   `json:"roleRef,omitempty"`
}

// Subject defines an identity. The identity is either a user or identified by a set of `properties`.
type Subject struct {
	User       string            `json:"user,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

// RoleRef refers to a role object.
type RoleRef struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
}

// ServiceRoleBindingStatus defines the observed state of ServiceRoleBinding
type ServiceRoleBindingStatus struct {
}

// ServiceRoleBinding is the Schema for the servicerolebindings API
type ServiceRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceRoleBindingSpec   `json:"spec,omitempty"`
	Status ServiceRoleBindingStatus `json:"status,omitempty"`
}

// ServiceRoleBindingList contains a list of ServiceRoleBinding
type ServiceRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceRoleBinding `json:"items"`
}

// GetBindingName TBU
// https://github.com/kubeflow/kubeflow/blob/d6bf8f85046fbc1ea5efa81cf0ef7905503dba8c/components/access-management/kfam/bindings.go#L58
func GetBindingName(userName string) (string, error) {
	// Only keep lower case letters and numbers, replace other with -
	reg, err := regexp.Compile("[^a-z0-9]+")
	if err != nil {
		return "", err
	}
	nameRaw := strings.ToLower(
		strings.Join([]string{
			"user",
			url.QueryEscape(reg.ReplaceAllString(userName, "-")),
			"clusterrole",
			"edit",
		}, "-"),
	)

	return reg.ReplaceAllString(nameRaw, "-"), nil
}
