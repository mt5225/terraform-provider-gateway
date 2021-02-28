package gateway

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gw "github.com/mt5225/cloudj-gateway-go/gateway"
	"github.com/tidwall/gjson"
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
		"access":         data.Get("access"),
		"head":           data.Get("head"),
		"interfaceeName": data.Get("interface_name"),
		"method":         data.Get("method"),
		"params":         data.Get("params"),
		"product":        data.Get("product"),
		"provider":       data.Get("rs_provider"),
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

	// invoke bindings to make request to gateway endpoint
	resp, err := gw.Create(opts)

	if err != nil {
		return fmt.Errorf("fail to parse json")
	}

	// we need to set the resource id before completely returning from this stack
	// for different integrations
	if data.Get("rs_provider") == "fusioncloud" {
		if resp.Success == false {
			return fmt.Errorf("fail to get server id from fusioncloud respone")
		}
		if resp.ServerID == "" {
			return fmt.Errorf("fail to get server id from fusioncloud respone")
		}
		data.SetId(resp.ServerID)
	}

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

func fusioncloudGetServerID(msg *map[string]interface{}) string {
	json := fmt.Sprintf("%v", msg)
	value := gjson.Get(json, "resultObject.resultMap.servers.0.tenant_id")
	return value.String()
}
