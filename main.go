package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/rvichery/terraform-provider-nuagenetworks/nuagenetworks"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nuagenetworks.Provider})
}
