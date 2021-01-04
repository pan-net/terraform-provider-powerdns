package powerdns

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccTestAccPDNSZoneDatasource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPDNSZoneDatasource,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.powerdns_zone.sysa", "kind", "Native"),
					resource.TestCheckResourceAttr("data.powerdns_zone.sysa", "soa_edit_api", "DEFAULT"),
				),
			},
		},
	})
}

const testAccPDNSZoneDatasource = `
data "powerdns_zone" "sysa" {
    name = "sysa.xyz."
}
`
