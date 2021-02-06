package admin

import (
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
)

// DeleteProfile delete kubeflow profile
func DeleteProfile(profileName string) {

	_, err := c.GetProfile(profileName)
	if err != nil {
		if errors.IsNotFound(err) {
			panic(fmt.Errorf("Kubeflow profile '%s' does not exist", profileName))

		} else {
			panic(err)
		}
	}

	err = c.DeleteProfile(profileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Profile '%s' deleted\n", profileName)
}
