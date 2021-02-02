package admin

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
)

// UpdateUser update kubeflow staticPassword's password
func UpdateUser(email, password string) {

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

	// Change Password
	hashedPassword, err := hashPassword(password)
	if err != nil {
		panic(err)
	}
	dc.StaticPasswords[userIdx].Hash = hashedPassword
	cm.Data["config.yaml"] = client.MarshalDexConfig(dc)

	// Restart
	err = c.UpdateConfigMap("auth", "dex", cm)
	if err != nil {
		panic(err)
	}

	err = c.RestartDexDeployment()
	if err != nil {
		fmt.Println("restart failed")
		fmt.Println(err)
		fmt.Println("rollback dex")

		cm = c.GetConfigMap("auth", "dex")
		cm.Data["config.yaml"] = originalData
		err := c.UpdateConfigMap("auth", "dex", cm)
		if err != nil {
			panic(err)
		}
		return
	}
	fmt.Printf("user '%s' updated\n", email)
}
