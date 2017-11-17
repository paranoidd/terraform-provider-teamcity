## 0.1.1 (Unreleased)

BACKWARDS INCOMPATIBILITIES / NOTES:

FEATURES:

- **Resource:**  `agent_pool_project_attachment` ([#10](https://github.com/Cardfree/terraform-provider-teamcity/pull/10]))

IMPROVEMENTS:

- **Resource:**  `build_configuration` and `build_template` Now support Settings ([#6](https://github.com/Cardfree/terraform-provider-teamcity/pull/6]))
- **Resource**: `vcs_root` Now support Propeties on Read ([#5](https://github.com/Cardfree/terraform-provider-teamcity/pull/5]))
- `parameter` Now supports `read_only = true/false` ([#7](https://github.com/Cardfree/terraform-provider-teamcity/pull/7]))
- `Step` Now ignores the constant `teamcity.step.mode` on read ([#8](https://github.com/Cardfree/terraform-provider-teamcity/pull/8/))

BUG FIXES:

- `parameter` and `parameter_value`replaces single element not all elements ([#9](https://github.com/Cardfree/terraform-provider-teamcity/pull/9/))
- Project Root is not defined in TF breaking compare ([#4](https://github.com/Cardfree/terraform-provider-teamcity/pull/4]))

## 0.1.0 (October 19, 2017)

BACKWARDS INCOMPATIBILITIES / NOTES:

FEATURES:

* **New Resource**: `vcs_root` ([#1](https://github.com/Cardfree/terraform-provider-teamcity/pull/1]))

IMPROVEMENTS:

* Support for Importing States on All Objects
* Added Website Docs

BUG FIXES:

- Vendored Go to keep changes at a minimum
