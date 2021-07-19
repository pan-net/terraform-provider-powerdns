Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) >=1.16 (to build the provider plugin)
-	[Goreleaser](https://goreleaser.com) >=v0.157.0 (for releasing provider plugin)

The Go ang Goreleaser minimum versions were set to be able to build plugin for Darwin/ARM64 architecture [see goreleaser notes.](https://goreleaser.com/deprecations/#builds-for-darwinarm64)

Using the Provider (TF 0.13+)
----------------------

```hcl
terraform {
  required_providers {
    powerdns = {
      source = "pan-net/powerdns"
    }
  }
}

provider "powerdns" {
  server_url = "https://host:port/"  # or use PDNS_SERVER_URL variable
  api_key    = "secret"              # or use PDNS_API_KEY variable
}
```

For detailed usage see [provider's documentation page](https://www.terraform.io/docs/providers/powerdns/index.html)

Building The Provider
---------------------

Clone the provider repository:

```sh
$ git clone git@github.com:terraform-providers/terraform-provider-powerdns
```

Navigate to repository directory:

```sh
$ cd terraform-provider-powerdns
```

Build repository:

```sh
$ go build
```

This will compile and place the provider binary, `terraform-provider-powerdns`, in the current directory.

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *recommended*).
You'll also need to have `$GOPATH/bin` in your `$PATH`.

Make sure the changes you performed pass linting:

```sh
$ make lint
```

To install the provider, run `make build`. This will build the provider and put the provider binary in the current working directory.

```sh
$ make build
```

In order to run local provider tests, you can simply run `make test`.

```sh
$ make test
```

For running acceptance tests locally, you'll need to use `docker-compose` to prepare the test environment:

```sh
docker-compose run --rm setup
```

After setup is done, run the acceptance tests with `make testacc` (note the env variables needed to interact with the PowerDNS container)

* HTTP

```sh
~$  PDNS_SERVER_URL=http://localhost:8081 \
    PDNS_API_KEY=secret \
    make testacc
````

* HTTPS

```sh
~$  PDNS_SERVER_URL=localhost:4443 \
    PDNS_API_KEY=secret \
    PDNS_CACERT=$(cat ./tests/files/ssl/rootCA/rootCA.crt) \
    make testacc
````


And finally cleanup containers spun up by `docker-compose`:

```sh
~$ docker-compose down
```
