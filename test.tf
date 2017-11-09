provider "teamcity" {
  url      = "http://localhost:8112"
  user     = "admin"
  password = "admin"
}

resource "teamcity_build_template" "terraform-provider-teamcity" {
  project = "Single"
  name    = "terraform-provider-teamcity"

  parameter {
    name = "env.MUH"
    type = "password"
  }

  parameter {
    name = "env.BLAH"
  }

  parameter_values = {
    "env.MUH"  = "Hello"
    "env.BLAH" = "123"
  }
}
