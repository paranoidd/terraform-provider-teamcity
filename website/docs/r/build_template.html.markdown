---
layout: "teamcity"
page_title: "Teamcity: build_template"
sidebar_current: "docs-teamcity-resource-build-template"
description: |-
  Provides a Teamcity Build Template resource. 
---

# teamcity\_build\_template

Provides a Teamcity Build Template resource. 

## Example Usage

```hcl
resource "teamcity_build_template" "default" {
  project = "DefaultProject"
  name    = "default-build-template"

  parameter {
    name = "env.TEAMCITY_PASSWORD"
    type = "password"
    allow_multiple  = false
  }
  
  parameter_values {
    env.TEST = "Hello"
    env.EXAMPLE  = "Hush Hush"
  }
}

resource "teamcity_build_configuration" "example" {
  name = "default-example"
  template = "${teamcity_build_template.far.id}"
  
  step {
    type = "simpleRunner"
    name = "first"

    properties = {
      script.content     = "npm install"
      teamcity.step.mode = "default"
      use.custom.script  = "true"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the Project.
* `project` - (Optional) ID of the Project the Build Template resides within.
    Defaults to `_Root`
* `description` - (Optional) Description of the Build Template.
* `template` - (Optional) template ID for the Build Template.
* `parameter` - (Optional) parameter/s defined for the Build Template.
* `parameter_values` - (Optional) parameter value/s defined for the Build Template.
* `step` - (Optional) Build step/s defined for the Build Template.
* `attached_vcs_root` - (Optional) VCS Root/s attached to the Build Template.

---

The `step` block supports:
* `type` - (Required) Type of the Step. Eg. `simpleRunner`, `MSBuild`, `jetbrains_powershell`, and etc.
* `name` - (Optional) Name of the Step.
* `properties` - (Optional) The Properties for the Build Step.

## Attributes Reference

The following attributes are exported:

* `id` - The Build Template ID
* `name` - The name of the Build Template
* `project` - ID of the Project the Build Template resides within.
* `description` - Description of the Build Template.
* `template` - (If Defined) template ID for the Build Template.
* `parameter` - (If Defined) parameter/s defined for the Build Template.
* `parameter_values` - (If Defined) parameter value/s defined for the Build Template.
* `step` - (If Defined) Build step/s defined for the Build Template.
* `attached_vcs_root` - (If Defined) VCS Root/s attached to the Build Template.

## Import

Build Templates can be imported using the `id`, e.g.

```
$ terraform import teamcity_build_template.foobar Root_DefaultBuildTemplate
```
