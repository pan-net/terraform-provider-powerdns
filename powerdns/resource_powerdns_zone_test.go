package powerdns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccPDNSZone(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPDNSZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPDNSZoneConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPDNSZoneExists("powerdns_zone.test"),
					resource.TestCheckResourceAttr("powerdns_zone.test", "name", "sysa.abc."),
					resource.TestCheckResourceAttr("powerdns_zone.test", "kind", "Native"),
				),
			},
		},
	})
}

func testAccCheckPDNSZoneDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "powerdns_zone" {
			continue
		}

		client := testAccProvider.Meta().(*Client)
		exists, err := client.ZoneExists(rs.Primary.Attributes["zone"])
		if err != nil {
			return fmt.Errorf("Error checking if zone still exists: %#v", rs.Primary.ID)
		}
		if exists {
			return fmt.Errorf("Zone still exists: %#v", rs.Primary.ID)
		}

	}
	return nil
}

func testAccCheckPDNSZoneExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		client := testAccProvider.Meta().(*Client)
		exists, err := client.ZoneExists(rs.Primary.Attributes["name"])
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("Zone does not exist: %#v", rs.Primary.ID)
		}
		return nil
	}
}

const testPDNSZoneConfig = `
resource "powerdns_zone" "test" {
	name = "sysa.abc."
	kind = "Native"
	nameservers = [ "ns1.sysa.abc.", "ns2.sysa.abc." ]
}`
