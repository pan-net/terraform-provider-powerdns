package powerdns

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataPDNSZone() *schema.Resource {
	return &schema.Resource{
		Read: dataPDNSZoneRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"kind": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"account": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"nameservers": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},

			"masters": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},

			"serial": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"soa": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"soa_edit_api": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataPDNSZoneRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	d.SetId(d.Get("name").(string))
	log.Printf("[DEBUG] Reading PowerDNS Zone: %s", d.Id())
	zoneInfo, err := client.GetZone(d.Id())
	if err != nil {
		return fmt.Errorf("Couldn't fetch PowerDNS Zone: %s", err)
	}

	d.Set("kind", zoneInfo.Kind)
	d.Set("soa_edit_api", zoneInfo.SoaEditAPI)
	d.Set("serial", zoneInfo.Serial)
	d.Set("account", zoneInfo.Account)

	if zoneInfo.Kind != "Slave" {
		nameservers, err := client.ListRecordsInRRSet(zoneInfo.Name, zoneInfo.Name, "NS")
		if err != nil {
			return fmt.Errorf("couldn't fetch zone %s nameservers from PowerDNS: %v", zoneInfo.Name, err)
		}

		var zoneNameservers []string
		for _, nameserver := range nameservers {
			zoneNameservers = append(zoneNameservers, nameserver.Content)
		}

		d.Set("nameservers", zoneNameservers)
	}

	soarecord, err := client.ListRecordsInRRSet(zoneInfo.Name, zoneInfo.Name, "SOA")
	if err != nil {
		return fmt.Errorf("couldn't fetch zone %s SOA from PowerDNS: %v", zoneInfo.Name, err)
	}

	var zoneSOA []string
	for _, soa := range soarecord {
		zoneSOA = append(zoneSOA, soa.Content)
	}

	zoneSOAStr := strings.Join(zoneSOA, "")

	d.Set("soa", zoneSOAStr)

	if strings.EqualFold(zoneInfo.Kind, "Slave") {
		d.Set("masters", zoneInfo.Masters)
	}

	return nil
}
