---
layout: "teamcity"
page_title: "Teamcity: vcs_root"
sidebar_current: "docs-teamcity-resource-vcs-root"
description: |-
  Provides a Teamcity VCS Root resource. 
---

# teamcity\_vcs\_root

Provides a Teamcity VCS Root resource. 

## Example Usage

```hcl
# Create a new vcs root
resource "teamcity_vcs_root" "default" {
  name = "default-vcs"

  properties = {
    url    = "https://github.com/Cardfree/teamcity-sdk-go"
    branch = "refs/heads/master"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the VCS Root.
* `properties` - (Required) The Properties for the VCS Root.
* `provider` - (Optional) The VCS Provider.
    This defaults to `jetbrains.git`. Other options include `svn` and `cvs`
* `project` - (Optional) The Project this should be created in.
    This defaults to `_Root`


## Attributes Reference

The following attributes are exported:

* `id` - The VCS Root ID
* `name` - The name of the VCS Root
* `project` - The project of the VCS Root


## Import

VCS Roots can be imported using the `id`, e.g.

```
$ terraform import teamcity_vcs_root.foobar Root_DefaultVcs
```
