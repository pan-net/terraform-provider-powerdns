package powerdns

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccPDNSZoneMetadata_Empty(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testPDNSZoneMetadataEmpty,
				ExpectError: regexp.MustCompile("'metadata' must not be empty"),
			},
		},
	})
}

func testAccCheckPDNSZoneMetadataDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "powerdns_zone_metadata" {
			continue
		}

		client := testAccProvider.Meta().(*Client)
		exists, err := client.RecordExistsByID(rs.Primary.Attributes["zone"], rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Error checking if record still exists: %#v", rs.Primary.ID)
		}
		if exists {
			return fmt.Errorf("Record still exists: %#v", rs.Primary.ID)
		}

	}
	return nil
}

func testAccCheckPDNSZoneMetadataExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		client := testAccProvider.Meta().(*Client)
		foundRecords, err := client.ListRecordsByID(rs.Primary.Attributes["zone"], rs.Primary.ID)
		if err != nil {
			return err
		}
		if len(foundRecords) == 0 {
			return fmt.Errorf("Record does not exist")
		}
		for _, rec := range foundRecords {
			if rec.ID() == rs.Primary.ID {
				return nil
			}
		}
		return fmt.Errorf("Record does not exist: %#v", rs.Primary.ID)
	}
}

const testPDNSZoneMetadataEmpty = `
resource "powerdns_zone_metadata" "test-a" {
	zone = "sysa.xyz."
	kind = "ALLOW-AXFR-FROM"
	metadata = [ ]
}`
