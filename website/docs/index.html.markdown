---
layout: "teamcity"
page_title: "Provider: Teamcity"
sidebar_current: "docs-teamcity-index"
description: |-
  The Teamcity provider is used to interact with Jetbrains Teamcity services. The provider needs to be configured with the proper credentials before it can be used.
---

# Teamcity Provider

The Teamcity provider is used to interact with
[Teamcity Server by Jetbrains](https://www.jetbrains.com/teamcity/). The provider needs
to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
// Configure the Teamcity provider
provider "teamcity" {
  url      = "${var.teamcity_url}"

  // username = "${var.teamcity_username}"
  // password = "${var.teamcity_password}"
}
```
## Teamcity versions
Compatibility is defined by the [teamcity-go-sdk](https://github.com/Cardfree/teamcity-go-sdk) which ships with this provider.
* Terraform `<= 0.9.6` - Teamcity `9.x+``


## Argument Reference

The following arguments are supported in the `provider` block:

* `url` - (optional) This is the Teamcity Server URL e.g. https://teamcity.domain.com:8111.
  It must be provided but it can also be sourced from the `TEAMCITY_URL` environment variable.
  Defaults to `http://localhost`

* `username` - (Optional) This is the Teamcity username. It must be provided, but
  it can also be sourced from the `TEAMCITY_USERNAME` environment variable.

* `password` - (Optional) This is the Teamcity Password. It must be provided, but
  it can also be sourced from the `TEAMCITY_PASSWORD` environment variable.

