---
layout: "teamcity"
page_title: "Teamcity: agent_pool"
sidebar_current: "docs-teamcity-datasource-agent-pool"
description: |-
  Provides details about a Teamcity Agent Pool.
---

# teamcity\_agent_pool

Provides details about a Teamcity Agent Pool. 

## Example Usage

```hcl
data "teamcity_agent_pool" "default" {
  name    = "Default"
}
```

## Argument Reference

The following arguments are supported:

* `Name` - (Required) Name of the Agent Pool.


## Attributes Reference

The following attributes are exported:

* `id` - The Agent Pool ID
* `Name` - The Agent Pool Name.

