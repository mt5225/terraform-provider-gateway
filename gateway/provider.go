package gateway

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

type GatewayConfiguration struct {
	endpoint string
}

// init provider block
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			// provider arguments and their specifications go here
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GW_ENDPOINT", nil),
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if value == "" {
						errors = append(errors, fmt.Errorf("Endpoint must not be an empty string"))
					}

					return
				},
			},
		},
		// map terraform dsl resources to functions
		ResourcesMap: map[string]*schema.Resource{
			"gateway_pack": resourcePack(),
		},
		// provider configuration function
		ConfigureFunc: configureProvider,
	}
}

// configure provider options
func configureProvider(data *schema.ResourceData) (interface{}, error) {
	// pass options from terraform DSL to the client
	endpoint := data.Get("endpoint").(string)

	// test endpoint
	err := connectGateway(endpoint)
	if err != nil {
		return nil, err
	}
	// code to error handle
	return &GatewayConfiguration{
		endpoint: endpoint,
	}, nil
}

func connectGateway(endpoint string) error {
	_, err := http.Get(endpoint)

	if err != nil {
		fmt.Errorf("Error connect to gateway ")
	}

	// more code for communicating the pizza request with the api

	return err
}
