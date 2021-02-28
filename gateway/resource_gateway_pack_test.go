package gateway

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccUser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testGatewayPackConfigbasic,
				Check: resource.ComposeTestCheckFunc(
					testEndpointExist("gateway_pack.test"),
					resource.TestCheckResourceAttr("gateway_pack.test", "method", "describeInstances"),
					resource.TestCheckResourceAttr("gateway_pack.test", "product", "ecs"),
					resource.TestCheckResourceAttr("gateway_pack.test", "provider", "fusionclouda"),
				),
			},
		},
	})
}

var testGatewayPackConfigbasic = `
resource "gateway_pack" "test" {
	access = {
		username = "testtenant"
        password = "Huawei@321"
        domainName = "testtenant"
        projectId = "c00fb95d354349d8b9f977ba319f2100"
        authEndpoint = "http://10.0.2.39:9102"
	}
	
	heads = {
    }

   interface_name = "Compute Interface"
   method = "describeInstances"
   params =  {
   }
   product = "ecs"
   provider = "fusioncloud"
}

output "resource_id" {
	value = "${gateway_pack.test.id}"
}
`

func testEndpointExist(rn string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource id not set")
		}
		return nil
	}
}
