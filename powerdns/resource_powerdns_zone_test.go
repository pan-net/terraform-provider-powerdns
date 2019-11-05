package powerdns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccPDNSZoneNative(t *testing.T) {
	resourceName := "powerdns_zone.test-native"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPDNSZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPDNSZoneConfigNative,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPDNSZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "sysa.abc."),
					resource.TestCheckResourceAttr(resourceName, "kind", "Native"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPDNSZoneMaster(t *testing.T) {
	resourceName := "powerdns_zone.test-master"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPDNSZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPDNSZoneConfigMaster,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPDNSZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "master.sysa.abc."),
					resource.TestCheckResourceAttr(resourceName, "kind", "Master"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPDNSZoneMasterSOAAPIEDIT(t *testing.T) {
	resourceName := "powerdns_zone.test-master-soa-edit-api"
	resourceSOAEDITAPI := `DEFAULT`

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPDNSZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPDNSZoneConfigMasterSOAEDITAPI,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPDNSZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "master-soa-edit-api.sysa.abc."),
					resource.TestCheckResourceAttr(resourceName, "kind", "Master"),
					resource.TestCheckResourceAttr(resourceName, "soa_edit_api", resourceSOAEDITAPI),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPDNSZoneMasterSOAAPIEDITEmpty(t *testing.T) {
	resourceName := "powerdns_zone.test-master-soa-edit-api-empty"
	resourceSOAEDITAPI := `""`

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPDNSZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPDNSZoneConfigMasterSOAEDITAPIEmpty,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPDNSZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "master-soa-edit-api-empty.sysa.abc."),
					resource.TestCheckResourceAttr(resourceName, "kind", "Master"),
					resource.TestCheckResourceAttr(resourceName, "soa_edit_api", resourceSOAEDITAPI),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPDNSZoneMasterSOAAPIEDITUndefined(t *testing.T) {
	resourceName := "powerdns_zone.test-master-soa-edit-api-undefined"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPDNSZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPDNSZoneConfigMasterSOAEDITAPIUndefined,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPDNSZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "master-soa-edit-api-undefined.sysa.abc."),
					resource.TestCheckResourceAttr(resourceName, "kind", "Master"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPDNSZoneSlave(t *testing.T) {
	resourceName := "powerdns_zone.test-slave"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPDNSZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPDNSZoneConfigSlave,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPDNSZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "slave.sysa.abc."),
					resource.TestCheckResourceAttr(resourceName, "kind", "Slave"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
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

const testPDNSZoneConfigNative = `
resource "powerdns_zone" "test-native" {
	name = "sysa.abc."
	kind = "Native"
	nameservers = ["ns1.sysa.abc.", "ns2.sysa.abc."]
}`

const testPDNSZoneConfigMaster = `
resource "powerdns_zone" "test-master" {
	name = "master.sysa.abc."
	kind = "Master"
	nameservers = ["ns1.sysa.abc.", "ns2.sysa.abc."]
}`

const testPDNSZoneConfigMasterSOAEDITAPI = `
resource "powerdns_zone" "test-master-soa-edit-api" {
	name = "master-soa-edit-api.sysa.abc."
	kind = "Master"
	nameservers = ["ns1.sysa.abc.", "ns2.sysa.abc."]
	soa_edit_api = "DEFAULT"
}`

const testPDNSZoneConfigMasterSOAEDITAPIEmpty = `
resource "powerdns_zone" "test-master-soa-edit-api-empty" {
	name = "master-soa-edit-api-empty.sysa.abc."
	kind = "Master"
	nameservers = ["ns1.sysa.abc.", "ns2.sysa.abc."]
	soa_edit_api = "\"\""
}`

const testPDNSZoneConfigMasterSOAEDITAPIUndefined = `
resource "powerdns_zone" "test-master-soa-edit-api-undefined" {
	name = "master-soa-edit-api-undefined.sysa.abc."
	kind = "Master"
	nameservers = ["ns1.sysa.abc.", "ns2.sysa.abc."]
}`

const testPDNSZoneConfigSlave = `
resource "powerdns_zone" "test-slave" {
	name = "slave.sysa.abc."
	kind = "Slave"
	nameservers = []
}`
