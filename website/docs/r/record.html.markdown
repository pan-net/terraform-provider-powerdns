---
layout: "powerdns"
page_title: "PowerDNS: powerdns_record"
sidebar_current: "docs-powerdns-resource-record"
description: |-
  Provides a PowerDNS record resource.
---

# powerdns\_record

Provides a PowerDNS record resource.

## Example Usage

Note that PowerDNS internally lowercases certain records (e.g. CNAME and AAAA), which can lead to resources being marked for a change in every singe plan.

For the v1 API (PowerDNS version 4):

```hcl
# Add a record to the zone
resource "powerdns_record" "foobar" {
  zone    = "example.com."
  name    = "www.example.com"
  type    = "A"
  ttl     = 300
  records = ["192.168.0.11"]
}
```

For PTR record example:
```hcl
# Add PTR record to the zone
resource "powerdns_record" "foobar" {
  zone    = "0.10.in-addr.arpa."
  name    = "10.0.0.10.in-addr.arpa."
  type    = "PTR"
  ttl     = 300
  records = ["www.example.com."]
}
```

For the legacy API (PowerDNS version 3.4):

```hcl
# Add a record to the zone
resource "powerdns_record" "foobar" {
  zone    = "example.com"
  name    = "www.example.com"
  type    = "A"
  ttl     = 300
  records = ["192.168.0.11"]
}
```

## Argument Reference

The following arguments are supported:

* `zone` - (Required) The name of zone to contain this record.
* `name` - (Required) The name of the record.
* `type` - (Required) The record type.
* `ttl` - (Required) The TTL of the record.
* `records` - (Required) A string list of records.

