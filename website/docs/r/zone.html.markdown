---
layout: "powerdns"
page_title: "PowerDNS: powerdns_zone"
sidebar_current: "docs-powerdns-zone"
description: |-
  Provides a PowerDNS zone.
---

# powerdns\_zone

Provides a PowerDNS zone.

## Example Usage

For the v1 API (PowerDNS version 4):

```hcl
# Add a zone
resource "powerdns_zone" "foobar" {
  name        = "example.com."
  kind        = "Native"
  nameservers = ["ns1.example.com.", "ns2.example.com."]
}
```

```hcl
# Add a Slave zone with list of IPs configured as a master for this zone
resource "powerdns_zone" "fubar" {
  name     = "slave.example.com."
  kind     = "Slave"
  masters  = ["10.10.10.10", "20.20.20.21"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of zone.
* `kind` - (Required) The kind of the zone.
* `nameservers` - (Optional) The zone nameservers.
* `masters` - (Optional) List of IP addresses configured as a master for this zone (“Slave” kind zones only)
* `soa_edit_api` - (Optional) This should map to one of the [supported API values](https://doc.powerdns.com/authoritative/dnsupdate.html#soa-edit-dnsupdate-settings) *or* in [case you wish to remove the setting](https://doc.powerdns.com/authoritative/domainmetadata.html#soa-edit-api), set this argument as `\"\"` (that will translate to the API value `""`).

## Importing

An existing zone can be imported into this resource by supplying the zone name. If the zone is not found, an error will be returned.

For example, to import zone `test.com.`:

```
$ terraform import powerdns_zone.test test.com.
```

For more information on how to use terraform's `import` command, please refer to terraform's [core documentation](https://www.terraform.io/docs/import/index.html#currently-state-only).
