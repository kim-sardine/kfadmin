package admin

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client/manifest"
)

// DeleteUser create kubeflow staticPassword
func DeleteUser(email string) {

	// Check if user exists
	cm, err := c.GetDex()
	if err != nil {
		panic(err)
	}

	originalData := cm.Data["config.yaml"]
	dc := manifest.UnmarshalDexConfig(originalData)
	users := dc.StaticPasswords
	userIdx := -1
	for i, user := range users {
		if user.Email == email {
			userIdx = i
			break
		}
	}
	if userIdx < 0 {
		panic(fmt.Errorf("user with email '%s' does not exist", email))
	}

	// Delete if exists
	dc.StaticPasswords = append(dc.StaticPasswords[:userIdx], dc.StaticPasswords[userIdx+1:]...)
	cm.Data["config.yaml"] = manifest.MarshalDexConfig(dc)
	err = c.UpdateDex(cm)
	if err != nil {
		panic(err)
	}

	err = c.RestartDexDeployment(originalData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user '%s' deleted\n", email)
}
