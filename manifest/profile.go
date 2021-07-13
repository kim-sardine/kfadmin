// https://github.com/kubeflow/kubeflow/blob/master/components/profile-controller/api/v1/profile_types.go

package manifest

import (
	"encoding/json"

	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Plugin is for customize actions on different platform.
type Plugin struct {
	metav1.TypeMeta `json:",inline"`
	Spec            *runtime.RawExtension `json:"spec,omitempty"`
}

// ProfileCondition TBU
type ProfileCondition struct {
	Type    string `json:"type,omitempty"`
	Status  string `json:"status,omitempty" description:"status of the condition, one of True, False, Unknown"`
	Message string `json:"message,omitempty"`
}

// ProfileSpec defines the desired state of Profile
type ProfileSpec struct {
	// The profile owner
	Owner   rbacv1.Subject `json:"owner,omitempty"`
	Plugins []Plugin       `json:"plugins,omitempty"`
	// Resourcequota that will be applied to target namespace
	ResourceQuotaSpec v1.ResourceQuotaSpec `json:"resourceQuotaSpec,omitempty"`
}

// ProfileStatus defines the observed state of Profile
type ProfileStatus struct {
	Conditions []ProfileCondition `json:"conditions,omitempty"`
}

// Profile is the Schema for the profiles API
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProfileSpec   `json:"spec,omitempty"`
	Status ProfileStatus `json:"status,omitempty"`
}

// Profiles contains a list of Profile
type Profiles struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

// GetProfile TBU
func GetProfile(profileName, email string) Profile {
	return Profile{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "kubeflow.org/v1",
			Kind:       "Profile",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: profileName,
		},
		Spec: ProfileSpec{
			Owner: rbacv1.Subject{
				Kind: "User",
				Name: email,
			},
		},
	}
}

// UnmarshalProfile TBU
func UnmarshalProfile(data []byte) (Profile, error) {
	var profile Profile
	err := json.Unmarshal(data, &profile)
	if err != nil {
		return Profile{}, err
	}
	return profile, nil
}

// UnmarshalProfiles TBU
func UnmarshalProfiles(data []byte) (Profiles, error) {
	var profiles Profiles
	err := json.Unmarshal(data, &profiles)
	if err != nil {
		return Profiles{}, err
	}
	return profiles, nil
}

// MarshalProfile TBU
func MarshalProfile(profile Profile) ([]byte, error) {
	data, err := json.Marshal(&profile)
	if err != nil {
		return nil, err
	}
	return data, nil
}
