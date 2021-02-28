package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/mt5225/terraform-provider-gateway/gateway"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gateway.Provider,
	})
}
