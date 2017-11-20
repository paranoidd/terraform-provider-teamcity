---
layout: "teamcity"
page_title: "Teamcity: agent_pool_project_attachment"
sidebar_current: "docs-teamcity-resource-agent-pool-project-attachment"
description: |-
  Provides a Teamcity Agent Pool Project Attachment resource. 
---

# teamcity\_agent_pool_project_attachment

Provides a Teamcity Agent Pool Project Attachment resource. 

Project/s can be attached to multiple Agent Pools.

## Example Usage

```hcl
data "teamcity_agent_pool" "default" {
  name    = "Default"
}

resource "teamcity_agent_pool_project_attachment" "default" {
  pool    = "${data.teamcity_agent_pool.default.id}"
  project = "${teamcity_project.default.id}"
}
```

## Argument Reference

The following arguments are supported:

* `pool` - (Required) ID of the Agent Pool.
* `project` - (Required) ID of the Project to Attach.


## Attributes Reference

The following attributes are exported:

* `id` - The Combination of Pool ID and Project ID
* `pool` -  The ID of the Agent Pool.
* `project` - The ID of the Project to Attach.

