package admin

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client/manifest"
)

// CreateUser create kubeflow staticPassword
func CreateUser(email, password string) {

	username, err := getUsernameFromEmail(email)
	if err != nil {
		panic(err)
	}

	cm, err := c.GetDex()
	if err != nil {
		panic(err)
	}

	originalData := cm.Data["config.yaml"]
	dc := manifest.UnmarshalDexConfig(originalData)
	users := dc.StaticPasswords

	uuids := make([]string, len(users)+1)
	for _, user := range users {
		uuids = append(uuids, user.UserID)
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		panic(err)
	}

	newUser := manifest.StaticPassword{
		Email:    email,
		Hash:     hashedPassword,
		Username: username,
		UserID:   getUniqueUUID(uuids),
	}

	dc.StaticPasswords = append(dc.StaticPasswords, newUser)

	cm.Data["config.yaml"] = manifest.MarshalDexConfig(dc)
	err = c.UpdateDex(cm)
	if err != nil {
		panic(err)
	}

	err = c.RestartDexDeployment(originalData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user '%s' created\n", email)
}
