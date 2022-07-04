package powerdns

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"net"
	"strings"
)

func expandIPv6Address(ip net.IP) string {
	b := make([]byte, 0, len(ip))

	// Print with possible :: in place of run of zeros
	for i := 0; i < len(ip); i += 2 {
		if i > 0 {
			b = append(b, ':')
		}
		s := (uint32(ip[i]) << 8) | uint32(ip[i+1])
		bHex := fmt.Sprintf("%04x", s)
		b = append(b, bHex...)
	}
	return string(b)
}

func resourcePDNSPTR() *schema.Resource {
	return &schema.Resource{
		Read: resourcePDNSPTRRead,
		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ptr_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourcePDNSPTRRead(d *schema.ResourceData, meta interface{}) error {
	ipAddressStr := d.Get("ip_address").(string)

	ipAddress := net.ParseIP(ipAddressStr)
	if ipAddress == nil {
		return fmt.Errorf("%v is not a valid IP address", ipAddressStr)
	}

	d.SetId(ipAddressStr)

	ipAddress4 := ipAddress.To4()

	if ipAddress4 != nil {
		// IPv4
		addressStringSplitted := strings.Split(ipAddress4.String(), ".")
		reverseAddressParts := make([]string, 0)
		for i := len(addressStringSplitted) - 1; i >= 0; i-- {
			reverseAddressParts = append(reverseAddressParts, addressStringSplitted[i])
		}
		reverseAddress := strings.Join(reverseAddressParts, ".")

		ptrRecord := fmt.Sprintf("%v.in-addr.arpa.", reverseAddress)
		return d.Set("ptr_address", ptrRecord)
	}

	expandedAddress := expandIPv6Address(ipAddress)

	addressStringSplitted := strings.Split(strings.ReplaceAll(expandedAddress, ":", ""), "")
	reverseAddressParts := []string{}
	for i := len(addressStringSplitted) - 1; i >= 0; i-- {
		reverseAddressParts = append(reverseAddressParts, addressStringSplitted[i])
	}
	reverseAddress := strings.Join(reverseAddressParts, ".")

	ptrRecord := fmt.Sprintf("%v.ip6.arpa.", reverseAddress)
	return d.Set("ptr_address", ptrRecord)

}
