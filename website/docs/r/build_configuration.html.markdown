---
layout: "teamcity"
page_title: "Teamcity: build_configuration"
sidebar_current: "docs-teamcity-resource-build-configuration"
description: |-
  Provides a Teamcity Build Configuration resource. 
---

# teamcity\_build\_configuration

Provides a Teamcity Build Configuration resource. 

## Example Usage

```hcl
resource "teamcity_build_configuration" "default" {
  project = "DefaultProject"
  name    = "default-build-configuration"

  setting {
     name  = "artifactRules"
     value = "+:dist/release.tar.gz => dist/"
  }

  parameter {
    name = "env.TEAMCITY_PASSWORD"
    type = "password"
    allow_multiple  = false
  }

  parameter {
    name            = "env.TEST"
    type            = "text"
    validation_mode = "not_empty"
    label           = "Test framework"
    description     = "Name of the test framework to use"
    allow_multiple  = false
  }

  parameter_values {
    "env.TEST" = "Hello"
  }

  step {
    type = "simpleRunner"
    name = "npm"

    properties = {
      script.content     = <<EOF
npm run
npm install
EOF

      teamcity.step.mode = "default"
      use.custom.script  = "true"
    }
  }

  step {
    type = "simpleRunner"
    name = "second"

    properties = {
      script.content     = "npm uninstall"
      teamcity.step.mode = "default"
      use.custom.script  = "true"
    }
  }

  attached_vcs_root {
    vcs_root       = "Root_DefaultVcs"
    checkout_rules = "+:. => .repo"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the Build Configuration.
* `project` - (Optional) ID of the Project the Build Configuration resides within.
    Defaults to `_Root`
* `description` - (Optional) Description of the Build Configuration.
* `setting` - (Optional) Build setting/s defined for the Build Configuration.
    - `name` - (Required) Name of the Setting.
    - `value` - (Optional) The Value of the Setting.
* `template` - (Optional) template ID for the Build Configuration.
* `parameter` - (Optional) parameter/s defined for the Build Configuration.
* `parameter_values` - (Optional) parameter value/s defined for the Build Configuration.
* `step` - (Optional) Build step/s defined for the Build Configuration.
    * `type` - (Required) Type of the Step. Eg. `simpleRunner`, `MSBuild`, `jetbrains_powershell`, and etc.
    * `name` - (Optional) Name of the Step.
    * `properties` - (Optional) The Properties for the Step.
* `feature`- (Optional)
    * `type` - (Required) Type of the Feature.
    * `properties` - (Optional) The Properties for the Feature.
* `trigger` - (Optional) a rule which initiates a new build on certain events.
    * `type` - (Required) Type of the Trigger.
    * `properties` - (Optional) The Properties for the Trigger.
* `snapshot_dependency`- (Optional) A build configuration can be made dependent on the sources of builds
    * `type` - (Required) Type of the Snapshot Dependency.
    * `dependent` (Required) - ID of the Build your are Dependent on.
    * `properties` - (Optional) The Properties for the Snapshot Dependency.
* `artifact_dependency` - (Optional) A build configuration can be made dependent on the artifacts 
    * `type` - (Required) Type of the Artifact Dependency.
    * `dependent` (Required) - ID of the Build your are Dependent on.
    * `properties` - (Optional) The Properties for the Artifact Dependency.
* `agent_requirement` - (Optional)  a piece of functionality that  affect running builds or reporting build results.
    * `name` - (Optional) Name of the Agent Requirement.
    * `properties` - (Optional) The Properties for the Agent Requirement.
* `attached_vcs_root` - (Optional) VCS Root/s attached to the Build Configuration.

## Attributes Reference

The following attributes are exported:

* `id` - The Build Configuration. ID
* `name` - The name of the Build Configuration.
* `project` - ID of the Project the Build Configuration resides within.
* `description` - Description of the Build Configuration.
* `setting` - (If Defined) Build setting/s defined for the Build Configuration.
* `template` - (If Defined) template ID for the Build Configuration.
* `parameter` - (If Defined) parameter/s defined for the Build Configuration.
* `parameter_values` - (If Defined) parameter value/s defined for the Build Configuration.
* `step` - (If Defined) Build step/s defined for the Build Configuration.
* `feature` - (If Defined) Build feature/s defined for the Build Configuration.
* `trigger` - (If Defined) Build trigger/s defined for the Build Configuration.
* `snapshot_dependency` - (If Defined) Build snapshot dependency/s defined for the Build Configuration.
* `artifact_dependency` - (If Defined) Build artifact dependency/s defined for the Build Configuration.
* `agent_requirement` - (If Defined) Build agent requirement/s defined for the Build Configuration.
* `attached_vcs_root` - (If Defined) VCS Root/s attached to the Build Configuration.

## Import

Build Configurations can be imported using the `id`, e.g.

```
$ terraform import teamcity_build_configuration.default Root_DefaultBuildConfiguration
```
