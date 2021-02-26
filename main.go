package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mt5225/terraform-provider-cloudj-gateway/gateway"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gateway.Provider,
	})
}