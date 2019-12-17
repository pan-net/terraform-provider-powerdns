package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-powerdns/powerdns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: powerdns.Provider})
}
