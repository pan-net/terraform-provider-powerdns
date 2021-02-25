---
layout: "powerdns"
page_title: "Provider: PowerDNS"
sidebar_current: "docs-powerdns-index"
description: |-
  The PowerDNS provider is used manipulate DNS records supported by PowerDNS server. The provider needs to be configured with the proper credentials before it can be used.
---

# PowerDNS Provider

The PowerDNS provider is used manipulate DNS records supported by PowerDNS server. The provider needs to be configured
with the proper credentials before it can be used. It supports both the [legacy API](https://doc.powerdns.com/3/httpapi/api_spec/) and the new [version 1 API](https://doc.powerdns.com/md/httpapi/api_spec/), however resources may need to be configured differently.

NOTE: if you're using the sqlite3 PowerDNS backend, you might face a problem (as described in [#75](https://github.com/pan-net/terraform-provider-powerdns/issues/75)) with terraform's
default behavior to [run mulitple operations](https://www.terraform.io/docs/commands/apply.html#parallelism-n) in parallel. Using `-parallelism=1` can help solve the limitations of
the sqlite3 PowerDNS Backend. The MySQL Backend has been verified to work with parallelism, however.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the PowerDNS provider
provider "powerdns" {
  api_key    = "${var.pdns_api_key}"
  server_url = "${var.pdns_server_url}"
}

# Create a record
resource "powerdns_record" "www" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `api_key` - (Required) The PowerDNS API key. This can also be specified with `PDNS_API_KEY` environment variable.
* `server_url` - (Required) The address of PowerDNS server. This can also be specified with `PDNS_SERVER_URL` environment variable. When no schema is provided, the default is `https`.
* `ca_certificate` - (Optional) A valid path of a Root CA Certificate in PEM format _or_ the content of a Root CA certificate in PEM format. This can also be specified with `PDNS_CACERT` environment variable.
* `insecure_https` - (Optional) Set this to `true` to disable verification of the PowerDNS server's TLS certificate. This can also be specified with the `PDNS_INSECURE_HTTPS` environment variable.
* `cache_requests` - (Optional) Set this to `true` to enable cache of the PowerDNS REST API requests. This can also be specified with the `PDNS_CACHE_REQUESTS` environment variable. `WARNING! Enabling this option can lead to the use of stale records when you use other automation to populate the DNS zone records at the same time.`
* `cache_mem_size` - (Optional) Memory size in MB for a cache of the PowerDNS REST API requests. This can also be specified with the `PDNS_CACHE_MEM_SIZE` environment variable.
* `cache_ttl` - (Optional) TTL in seconds for a cache of the PowerDNS REST API requests. This can also be specified with the `PDNS_CACHE_TTL` environment variable.
