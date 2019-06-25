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
  name    = "example.com."
  kind    = "Native"
  nameservers = ["ns1.example.com.", "ns2.example.com."]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of zone.
* `kind` - (Required) The kind of the zone.
* `nameservers` - (Required) The zone nameservers.
