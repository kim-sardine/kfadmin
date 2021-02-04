package admin

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client/manifest"
)

// UpdateUserPassword update kubeflow staticPassword's password
func UpdateUserPassword(email, password string) {

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

	// Change Password
	hashedPassword, err := hashPassword(password)
	if err != nil {
		panic(err)
	}

	dc.StaticPasswords[userIdx].Hash = hashedPassword
	cm.Data["config.yaml"] = manifest.MarshalDexConfig(dc)
	err = c.UpdateDex(cm)
	if err != nil {
		panic(err)
	}

	err = c.RestartDexDeployment(originalData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user '%s' updated\n", email)
}
