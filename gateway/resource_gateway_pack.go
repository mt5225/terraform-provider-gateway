package gateway

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gw "github.com/mt5225/cloudj-gateway-go/gateway"
)

// pizza resource declaration and schema
func resourcePack() *schema.Resource {
	return &schema.Resource{
		// functions for the various actions
		Create: resourcePackCreate,
		// Read:   resourcePackRead,
		// Update: resourcePackUpdate,
		// Delete: resourcePackDelete,
		// Exists: resourcePackExists,

		Schema: map[string]*schema.Schema{
			// resource arguments and their specifications go here
			"access": &schema.Schema{
				Type:        schema.TypeMap,
				Required:    true,
				Description: "the access info to resource",
			},
			"heads": &schema.Schema{
				Type:        schema.TypeMap,
				Required:    true,
				Description: "request header",
			},
			"interfaceName": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "interface name",
			},
			"method": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "request method",
			},
			"params": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "request parameters",
			},
			"product": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "product name",
			},
			"provider": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "provider name",
			},
		},
	}
}

// create a resource
func resourcePackCreate(data *schema.ResourceData, meta interface{}) error {
	// create struct from desired pizza arguments
	body := &map[string]interface{}{
		"access":        data.Get("access"),
		"head":          data.Get("head"),
		"interfaceName": data.Get("interfaceName"),
		"method":        data.Get("method"),
		"params":        data.Get("params"),
		"product":       data.Get("product"),
		"provider":      data.Get("provider"),
	}

	body_encode, _ := json.Marshal(*body)

	opts := map[string]interface{}{
		"endpoint": "",
		"method":   "POST",
		"body":     body_encode,
	}

	// invoke bindings to make pizza according to specifications
	resource, err := gw.Create(opts)

	// code to handle errors

	// we need to set the resource id before completely returning from this stack
	data.SetID(pizza["id"])

	return resourcePizzaRead(data, meta)
}
