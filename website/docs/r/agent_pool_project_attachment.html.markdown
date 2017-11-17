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
resource "teamcity_agent_pool_project_attachment" "default" {
  pool    = "Default"
  project = "${teamcity_project.default.id}"
}
```

## Argument Reference

The following arguments are supported:

* `pool` - (Required) Name of the Agent Pool.
* `project` - (Required) The ID of the Project to Attach.


## Attributes Reference

The following attributes are exported:

* `id` - The Combination of Pool ID and Project ID
* `pool` - (Required) Name of the Agent Pool.
* `project` - (Required) The ID of the Project to Attach.

