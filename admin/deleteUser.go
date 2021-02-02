package admin

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
)

// DeleteUser create kubeflow staticPassword
func DeleteUser(email string) {

	// Check if user exists
	cm := c.GetConfigMap("auth", "dex")
	originalData := cm.Data["config.yaml"]
	dc := client.UnmarshalDexConfig(originalData)
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

	// Restart
	err := c.UpdateConfigMap("auth", "dex", dc)
	if err != nil {
		panic(err)
	}

	err = c.RestartDexDeployment(originalData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user '%s' deleted\n", email)
}
