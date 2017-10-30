resource "teamcity_build_configuration" "import" {
  name    = "Import"
  project = "Empty"
}

resource "teamcity_project" "default" {
  name = "default-project"

  parameter {
    name = "env.READ_ONLY"
    type = "text"
  }
}
