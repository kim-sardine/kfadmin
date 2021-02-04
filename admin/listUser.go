package admin

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client/manifest"
)

// ListUser print staticPassword
func ListUser() {

	cm, err := c.GetDex()
	if err != nil {
		panic(err)
	}

	dc := manifest.UnmarshalDexConfig(cm.Data["config.yaml"])
	for i, user := range dc.StaticPasswords {
		fmt.Println(i+1, user.Email)
	}
	fmt.Printf("Total : %d users\n", len(dc.StaticPasswords))
}
