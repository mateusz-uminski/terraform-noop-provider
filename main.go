package main

import (
	"github.com/mateusz-uminski/terraform-providers/noop"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: noop.Provider,
	})
}
