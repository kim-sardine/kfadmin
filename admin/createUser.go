package admin

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
)

// CreateUser create kubeflow static password user
func CreateUser(email, password string) {

	username, err := getUsernameFromEmail(email)
	if err != nil {
		panic(err)
	}

	cm := c.GetConfigMap("auth", "dex")
	originalData := cm.Data["config.yaml"]
	dc := client.UnmarshalDexConfig(originalData)
	users := dc.StaticPasswords

	uuids := make([]string, len(users)+1)
	for _, user := range users {
		uuids = append(uuids, user.UserID)
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		panic(err)
	}

	newUser := client.StaticPasswordManifest{
		Email:    email,
		Hash:     hashedPassword,
		Username: username,
		UserID:   getUniqueUUID(uuids),
	}

	dc.StaticPasswords = append(dc.StaticPasswords, newUser)
	cm.Data["config.yaml"] = client.MarshalDexConfig(dc)

	err = c.UpdateConfigMap("auth", "dex", cm)
	if err != nil {
		panic(err)
	}

	err = c.RestartDexDeployment()
	if err != nil {
		// FIXME: Duplicate code
		fmt.Println("restart failed")
		fmt.Println(err)
		fmt.Println("rollback dex")

		cm = c.GetConfigMap("auth", "dex")
		cm.Data["config.yaml"] = originalData
		err := c.UpdateConfigMap("auth", "dex", cm)
		if err != nil {
			panic(err)
		}
		fmt.Println("user creation failed.")
		return
	}
	fmt.Printf("user '%s' created\n", email)
}
