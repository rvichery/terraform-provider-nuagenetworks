Nuage Networks Terraform Provider
==================
[![GoDoc](https://godoc.org/github.com/rvichery/terraform-provider-nuagenetworks/nuagenetworks?status.svg)](https://godoc.org/github.com/rvichery/terraform-provider-nuagenetworks/nuagenetworks)
[![Build Status](https://travis-ci.org/rvichery/terraform-provider-nuagenetworks.svg?branch=master)](https://travis-ci.org/rvichery/terraform-provider-nuagenetworks)
[![Go Report Card](https://goreportcard.com/badge/github.com/rvichery/terraform-provider-nuagenetworks)](https://goreportcard.com/report/github.com/rvichery/terraform-provider-nuagenetworks)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Frvichery%2Fterraform-provider-nuagenetworks.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Frvichery%2Fterraform-provider-nuagenetworks?ref=badge_shield)

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/rvichery/terraform-provider-nuagenetworks`

```sh
$ mkdir -p $GOPATH/src/github.com/rvichery/; cd $GOPATH/src/github.com/rvichery
$ git clone git@github.com:rvichery/terraform-provider-nuagenetworks
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/rvichery/terraform-provider-nuagenetworks
$ make build
```

Using the provider
----------------------
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

