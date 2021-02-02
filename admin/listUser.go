package admin

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
)

// ListUser print static users
func ListUser() {

	cm := c.GetConfigMap("auth", "dex")
	dc := client.UnmarshalDexConfig(cm.Data["config.yaml"])
	for i, user := range dc.StaticPasswords {
		fmt.Println(i+1, user.Email)
	}
	fmt.Printf("Total : %d users\n", len(dc.StaticPasswords))
}
