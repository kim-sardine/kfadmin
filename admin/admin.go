package admin

import (
	"github.com/kim-sardine/kfadmin/client"
)

var c *client.KfClient

func init() {
	c = &client.KfClient{}
	c.LoadClientset()
}
