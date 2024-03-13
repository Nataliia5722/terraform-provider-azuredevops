package main

import (
	"flag"

	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		Debug:        debug,
		ProviderAddr: "registry.terraform.io/microsoft/azuredevops",
		ProviderFunc: func() *schema.Provider {
			return azuredevops.Provider()
		},
	})
}
