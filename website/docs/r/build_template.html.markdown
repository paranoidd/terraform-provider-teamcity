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

* ## Argument Reference

  The following arguments are supported:

  - `name` - (Required) Name of the Build Configuration.
  - `project` - (Optional) ID of the Project the Build Configuration resides within.
    Defaults to `_Root`
  - `description` - (Optional) Description of the Build Configuration.
  - `parameter` - (Optional) parameter/s defined for the Build Configuration.
  - `parameter_values` - (Optional) parameter value/s defined for the Build Configuration.
  - `step` - (Optional) Build step/s defined for the Build Configuration.
    - `type` - (Required) Type of the Step. Eg. `simpleRunner`, `MSBuild`, `jetbrains_powershell`, and etc.
    - `name` - (Optional) Name of the Step.
    - `properties` - (Optional) The Properties for the Step.
  - `feature`- (Optional)
    - `type` - (Required) Type of the Feature.
    - `properties` - (Optional) The Properties for the Feature.
  - `trigger` - (Optional) a rule which initiates a new build on certain events.
    - `type` - (Required) Type of the Trigger.
    - `properties` - (Optional) The Properties for the Trigger.
  - `snapshot_dependency`- (Optional) A build configuration can be made dependent on the sources of builds
    - `type` - (Required) Type of the Snapshot Dependency.
    - `dependent` (Required) - ID of the Build your are Dependent on.
    - `properties` - (Optional) The Properties for the Snapshot Dependency.
  - `artifact_dependency` - (Optional) A build configuration can be made dependent on the artifacts 
    - `type` - (Required) Type of the Artifact Dependency.
    - `dependent` (Required) - ID of the Build your are Dependent on.
    - `properties` - (Optional) The Properties for the Artifact Dependency.
  - `agent_requirement` - (Optional)  a piece of functionality that  affect running builds or reporting build results.
    - `name` - (Optional) Name of the Agent Requirement.
    - `properties` - (Optional) The Properties for the Agent Requirement.
  - `attached_vcs_root` - (Optional) VCS Root/s attached to the Build Configuration.

  ## Attributes Reference

  The following attributes are exported:

  - `id` - The Build Configuration. ID
  - `name` - The name of the Build Configuration.
  - `project` - ID of the Project the Build Configuration resides within.
  - `description` - Description of the Build Configuration.
  - `template` - (If Defined) template ID for the Build Configuration.
  - `parameter` - (If Defined) parameter/s defined for the Build Configuration.
  - `parameter_values` - (If Defined) parameter value/s defined for the Build Configuration.
  - `step` - (If Defined) Build step/s defined for the Build Configuration.
  - `feature` - (If Defined) Build feature/s defined for the Build Configuration.
  - `trigger` - (If Defined) Build trigger/s defined for the Build Configuration.
  - `snapshot_dependency` - (If Defined) Build snapshot dependency/s defined for the Build Configuration.
  - `artifact_dependency` - (If Defined) Build artifact dependency/s defined for the Build Configuration.
  - `agent_requirement` - (If Defined) Build agent requirement/s defined for the Build Configuration.
  - `attached_vcs_root` - (If Defined) VCS Root/s attached to the Build Configuration.

## Import

Build Templates can be imported using the `id`, e.g.

```
$ terraform import teamcity_build_template.foobar Root_DefaultBuildTemplate
```
