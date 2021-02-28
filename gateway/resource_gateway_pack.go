package gateway

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gw "github.com/mt5225/cloudj-gateway-go/gateway"
)

// pizza resource declaration and schema
func resourcePack() *schema.Resource {
	return &schema.Resource{
		// functions for the various actions
		Create: resourcePackCreate,
		Read:   resourcePackRead,
		Update: resourcePackUpdate,
		Delete: resourcePackDelete,

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
			"interface_name": &schema.Schema{
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
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "request parameters",
			},
			"product": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "product name",
			},
			"rs_provider": &schema.Schema{
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
		"access":          data.Get("access"),
		"head":            data.Get("head"),
		"interfacee_name": data.Get("interface_name"),
		"method":          data.Get("method"),
		"params":          data.Get("params"),
		"product":         data.Get("product"),
		"rs_provider":     data.Get("provider"),
	}

	bodyEncode, err := json.Marshal(*body)
	if err != nil {
		log.Fatal("fail to marshal message body")
		return err
	}

	conf := meta.(*Configuration)
	opts := new(gw.Opts)
	opts.Endpoint = conf.endpoint
	opts.Method = "POST"
	opts.Body = bodyEncode

	// invoke bindings to make pizza according to specifications
	resource, err := gw.Create(opts)

	fmt.Print(resource)

	// code to handle errors

	// we need to set the resource id before completely returning from this stack
	data.SetId("abc")

	return nil
}

func resourcePackUpdate(data *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourcePackRead(data *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourcePackDelete(data *schema.ResourceData, meta interface{}) error {
	return nil
}
