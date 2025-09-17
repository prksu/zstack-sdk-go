package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

func (cli *ZSClient) QuerySshKeyPair(params param.QueryParam) ([]view.SshKeyInventoryView, error) {
	var resp []view.SshKeyInventoryView
	return resp, cli.List("v1/ssh-key-pair", &params, &resp)
}
