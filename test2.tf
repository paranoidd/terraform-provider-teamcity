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

  parameter {
    allow_multiple = "false"
    name           = "secure-string"
    type           = "password"
  }

  parameter_values {
    foo = "bar"
  }

  step {
    type = "simpleRunner"
    name = "second"

    properties = {
      script.content    = "echo 'Terraform'"
      use.custom.script = "true"
    }
  }
}
