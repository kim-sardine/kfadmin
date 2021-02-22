package manifest

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

var (
	testMarshaledProfile = "{\"kind\":\"Profile\",\"apiVersion\":\"kubeflow.org/v1\",\"metadata\":{\"name\":\"test-ns\",\"creationTimestamp\":null},\"spec\":{\"owner\":{\"kind\":\"User\",\"name\":\"tester@test.test\"},\"resourceQuotaSpec\":{}},\"status\":{}}"
)

func TestGetProfile(t *testing.T) {
	profile := GetProfile(testNamespace, testUserName)

	assert.Equal(t, profile.APIVersion, "kubeflow.org/v1")
	assert.Equal(t, profile.Kind, "Profile")
	assert.Equal(t, profile.ObjectMeta.Name, testNamespace)
	assert.Equal(t, profile.Spec.Owner.Kind, "User")
	assert.Equal(t, profile.Spec.Owner.Name, testUserName)
}

func TestMarshalProfile(t *testing.T) {
	profile := GetProfile(testNamespace, testUserName)

	data, _ := MarshalProfile(profile)

	assert.Equal(t, string(data), testMarshaledProfile)
}

func TestUnmarshalProfile(t *testing.T) {
	profile := GetProfile(testNamespace, testUserName)

	unmarshaledProfile, _ := UnmarshalProfile([]byte(testMarshaledProfile))

	assert.Equal(t, cmp.Equal(profile, unmarshaledProfile), true)
}
