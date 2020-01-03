package powerdns

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePDNSZone() *schema.Resource {
	return &schema.Resource{
		Create: resourcePDNSZoneCreate,
		Read:   resourcePDNSZoneRead,
		Update: resourcePDNSZoneUpdate,
		Delete: resourcePDNSZoneDelete,
		Exists: resourcePDNSZoneExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"kind": {
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
			},

			"nameservers": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"soa_edit_api": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
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
		SoaEditAPI:  d.Get("soa_edit_api").(string),
	}

	createdZoneInfo, err := client.CreateZone(zoneInfo)
	if err != nil {
		return err
	}

	d.SetId(createdZoneInfo.ID)

	return nil
}

func resourcePDNSZoneRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	log.Printf("[DEBUG] Reading PowerDNS Zone: %s", d.Id())
	zoneInfo, err := client.GetZone(d.Id())
	if err != nil {
		return fmt.Errorf("Couldn't fetch PowerDNS Zone: %s", err)
	}

	d.Set("name", zoneInfo.Name)
	d.Set("kind", zoneInfo.Kind)
	d.Set("soa_edit_api", zoneInfo.SoaEditAPI)

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

	return nil
}

func resourcePDNSZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Updating PowerDNS Zone: %s", d.Id())

	client := meta.(*Client)

	zoneInfo := ZoneInfo{}
	shouldUpdate := false
	if d.HasChange("kind") {
		zoneInfo.Kind = d.Get("kind").(string)
		shouldUpdate = true
	}

	if shouldUpdate {
		return client.UpdateZone(d.Id(), zoneInfo)
	}
	return nil
}

func resourcePDNSZoneDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	log.Printf("[INFO] Deleting PowerDNS Zone: %s", d.Id())
	err := client.DeleteZone(d.Id())

	if err != nil {
		return fmt.Errorf("Error deleting PowerDNS Zone: %s", err)
	}
	return nil
}

func resourcePDNSZoneExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Printf("[INFO] Checking existence of PowerDNS Zone: %s", d.Id())

	client := meta.(*Client)
	exists, err := client.ZoneExists(d.Id())

	if err != nil {
		return false, fmt.Errorf("Error checking PowerDNS Zone: %s", err)
	}
	return exists, nil
}
