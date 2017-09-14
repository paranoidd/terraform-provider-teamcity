package main

import (
	"github.com/Cardfree/terraform-provider-teamcity/teamcity"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: teamcity.Provider})
}
