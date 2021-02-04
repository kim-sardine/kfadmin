package admin

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client/manifest"
	"k8s.io/apimachinery/pkg/api/errors"
)

// CreateProfile create kubeflow profile
func CreateProfile(profileName, email string) {

	_, err := c.GetProfile(profileName)
	if err == nil {
		panic(fmt.Errorf("Profile '%s' already exists", profileName))
	}
	if !errors.IsNotFound(err) {
		panic(err)
	}

	users, err := c.GetStaticUsers()
	if err != nil {
		panic(err)
	}
	userExists := false
	for _, user := range users {
		if user.Email == email {
			userExists = true
			break
		}
	}
	if !userExists {
		panic(fmt.Errorf("User with email '%s' does not exist", email))
	}

	profile := manifest.GetProfile(profileName, email)
	err = c.CreateProfile(profile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Profile '%s' created\n", profileName)
}
