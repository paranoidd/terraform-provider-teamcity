resource "teamcity_project" "default" {
  name = "default-project"

  parameter {
    name = "env.READ_ONLY"
    type = "text"

    read_only = true

    # validation_mode = "any"
  }
}
