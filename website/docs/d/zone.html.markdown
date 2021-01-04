---
layout: "powerdns"
page_title: "PowerDNS: powerdns_zone"
sidebar_current: "docs-powerdns-datasource-zone"
description: |-
  Get information on a PowerDNS zone.
---

# powerdns\_zone

Use this data source to get informations on a PowerDNS zone.

## Example Usage

For the v1 API (PowerDNS version 4):

```hcl
data "powerdns_zone" "foobar" {
  name        = "example.com."
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of zone.

## Attributes Reference

The following attributes are exported:

* `name` - The name of zone.
* `id` - The PowerDNS id for the zone
* `kind` - The kind of the zone.
* `account` - The name of the account owning the zone.
* `nameservers` - The zone nameservers.
* `masters` - List of IP addresses configured as a master for this zone (“Slave” kind zones only).
* `serial` - The zone serial.
* `soa` - The zone SOA.
* `soa_edit_api` - The SOA edit API mode for the zone (https://doc.powerdns.com/authoritative/dnsupdate.html#soa-edit-dnsupdate-settings).