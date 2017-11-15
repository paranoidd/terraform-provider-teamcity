resource "teamcity_project" "default" {
  name = "default-project"

  parameter {
    name = "env.READ_ONLY"
    type = "text"

    read_only = true

    # validation_mode = "any"
  }
}

resource "teamcity_build_configuration" "default" {
  project = "${teamcity_project.default.id}"
  name    = "default-build-configuration"

  step {
    type = "simpleRunner"
    name = "second"

    properties = {
      script.content    = "echo 'Terraform'"
      use.custom.script = "true"
    }
  }
}
