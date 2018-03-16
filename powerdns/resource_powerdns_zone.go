package powerdns

import (
	"log"

	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePDNSZone() *schema.Resource {
	return &schema.Resource{
		Create: resourcePDNSZoneCreate,
		Read:   resourcePDNSZoneRead,
		Delete: resourcePDNSZoneDelete,
		Exists: resourcePDNSZoneExists,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"kind": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"nameservers": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePDNSZoneCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	var nameservers []string
	for _, nameserver := range d.Get("nameservers").(*schema.Set).List() {
		nameservers = append(nameservers, nameserver.(string))
	}

	zoneInfo := ZoneInfo{
		Name:        d.Get("name").(string),
		Kind:        d.Get("kind").(string),
		Nameservers: nameservers,
	}

	if err := client.CreateZone(zoneInfo); err != nil {
		return err
	}

	d.SetId(zoneInfo.Name)

	return nil
}

func resourcePDNSZoneRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	log.Printf("[DEBUG] Reading PowerDNS Zone: %s", d.Id())
	zoneInfo, err := client.GetZone(d.Get("name").(string))
	if err != nil {
		return fmt.Errorf("Couldn't fetch PowerDNS Zone: %s", err)
	}

	d.Set("name", zoneInfo.Name)
	d.Set("kind", zoneInfo.Kind)

	return nil
}

func resourcePDNSZoneDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	log.Printf("[INFO] Deleting PowerDNS Zone: %s", d.Id())
	err := client.DeleteZone(d.Get("name").(string))

	if err != nil {
		return fmt.Errorf("Error deleting PowerDNS Zone: %s", err)
	}

	return nil
}

func resourcePDNSZoneExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	name := d.Get("name").(string)

	log.Printf("[INFO] Checking existence of PowerDNS Zone: %s", name)

	client := meta.(*Client)
	exists, err := client.ZoneExists(name)

	if err != nil {
		return false, fmt.Errorf("Error checking PowerDNS Zone: %s", err)
	}
	return exists, nil
}
