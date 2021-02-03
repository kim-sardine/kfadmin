package admin

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client/manifest"
	"k8s.io/apimachinery/pkg/api/errors"
)

// CreateProfile create kubeflow profile
func CreateProfile(profileName, email string) {

	// check profile exists
	// check user exists
	// create profile if profile not exist and user exists

	_, err := c.GetProfile(profileName)
	// if profile exists
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

	profile := manifest.CreateProfileManifest(profileName, email)
	err = c.CreateProfile(profile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Profile '%s' created\n", profileName)
}
