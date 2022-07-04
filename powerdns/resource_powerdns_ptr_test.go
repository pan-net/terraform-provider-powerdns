package powerdns

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccPDNSPTR_v4(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testPDNSPtrConfig4,
				Check:  resource.TestCheckResourceAttr("data.powerdns_ptr.test4", "ptr_address", "4.3.2.1.in-addr.arpa."),
			},
		},
	})
}

func TestAccPDNSPTR_v6(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testPDNSPtrConfig6,
				Check:  resource.TestCheckResourceAttr("data.powerdns_ptr.test6", "ptr_address", "1.1.b.6.f.a.d.1.5.e.6.7.c.0.6.d.5.9.2.0.0.c.a.8.8.0.1.8.2.0.a.2.ip6.arpa."),
			},
		},
	})
}

const testPDNSPtrConfig4 = `
data "powerdns_ptr" "test4" {
	ip_address = "1.2.3.4"
}`

const testPDNSPtrConfig6 = `
data "powerdns_ptr" "test6" {
	ip_address = "2a02:8108:8ac0:295:d60c:76e5:1daf:6b11"
}`
