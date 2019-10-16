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
  name    = "www.example.com."
  type    = "A"
  ttl     = 300
  records = ["192.168.0.11"]
}
```

Here is an example creating PTR record:

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

PowerDNS API offers the feature to automatically create corresponding PTR record for the A/AAAA record.
Existing PTR records with same name are replaced. If no matching reverse zone is found, resource creation will fail. 
You can use `powerdns_zone` resource to create the reverse zone.
Here is an example of creating A record along with corresponding PTR record:

```hcl
resource "powerdns_record" "foobar" {
  zone    = "example.com."
  name    = "www.example.com"
  type    = "A"
  ttl     = 300
  set_ptr = true
  records = ["192.168.0.11"]
}
```

For the legacy API (PowerDNS version 3.4):

```hcl
# Add a record to the zone
resource "powerdns_record" "foobar" {
  zone    = "example.com."
  name    = "www.example.com."
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

### Attribute Reference

The id of the resource is a composite of the record name and record type, joined by a separator - `:::`.

For example, record `foo.test.com.` of type `A` will be represented with the following `id`: `foo.test.com.:::A`

### Importing

An existing record can be imported into this resource by supplying both the record id and zone name it belongs to.
If the record or zone is not found, or if the record is of a different type or in a different zone, an error will be returned.

For example:

```
$ terraform import powerdns_record.test-a '{"zone": "test.com.", "id": "foo.test.com.:::A"}'
```

For more information on how to use terraform's `import` command, please refer to terraform's [core documentation](https://www.terraform.io/docs/import/index.html#currently-state-only).

