---
layout: "teamcity"
page_title: "TeamCity: project"
sidebar_current: "docs-teamcity-resource-project"
description: |-
  Provides a TeamCity Project resource.
---

# teamcity\_project

Provides a TeamCity Project resource.

## Example Usage

```hcl
resource "teamcity_project" "default" {
  parent = "_Root"
  name   = "default-project"

  parameter {
    name = "env.TEAMCITY_PASSWORD"
    type = "password"
  }

  parameter {
    name            = "env.TEST"
    type            = "text"
    validation_mode = "not_empty"
    label           = "Test framework"
    description     = "Name of the test framework to use"
  }
  
  parameter_values {
    env.TEST = "Hello"
    env.EXAMPLE  = "Hush Hush"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the Project.
* `project` - (Optional) ID of the Project the Project resides within.
    Defaults to `_Root`
* `description` - (Optional) Description of the Project.
* `parameter` - (Optional) parameter/s defined for the Project.
* `parameter_values` - (Optional) parameter value/s defined for the Project.


## Attributes Reference

The following attributes are exported:

* `id` - The Project ID
* `name` - The name of the Project
* `project` - ID of the Project the Project resides within.
* `description` - Description of the Project.


## Import

Projects can be imported using the `id`, e.g.

```
$ terraform import teamcity_project.default DefaultProject
```
