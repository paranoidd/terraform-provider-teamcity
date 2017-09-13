package main

import (
	"fmt"
	//"log"
	//"regexp"
	//"strings"

	//"math/rand"
	// "reflect"
	// "testing"
	//"time"

	//"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/Cardfree/teamcity-sdk-go/teamcity"
	"github.com/Cardfree/teamcity-sdk-go/types"
)

var testAccVcsRoot = `
resource "teamcity_vcs_root" "bar" {
  project = "Single"
  name = "bar"
}`

// func TestAccVcsRoot_basic(t *testing.T) {
// 	var v types.VcsRoot

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckProjectDestroy,
// 		Steps: []resource.TestStep{
// 			resource.TestStep{
// 				Config: testAccProject,
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckProjectExists("teamcity_project.bar", &v),
// 					testAccCheckProjectAttributes(&v),
// 					testAccCheckProjectParameters(&v, &types.Parameters{
// 						"env.TEST": types.Parameter{
// 							Value: "Hello",
// 							Spec: &types.ParameterSpec{
// 								Label:       "Test framework",
// 								Description: "Name of the test framework to use",
// 								Type:        types.TextType{"not_empty"},
// 							},
// 						},
// 						"env.MUH": types.Parameter{
// 							Spec: &types.ParameterSpec{
// 								Type: types.PasswordType{},
// 							},
// 						},
// 					}),
// 					resource.TestCheckResourceAttr(
// 						"teamcity_project.bar", "parent", "Single"),
// 					resource.TestCheckResourceAttr(
// 						"teamcity_project.bar", "name", "bar"),
// 					resource.TestCheckResourceAttr(
// 						"teamcity_project.bar", "description", ""),
// 					resource.TestCheckResourceAttr(
// 						"teamcity_project.bar", "parameter_values.env.TEST", "Hello"),
// 					testAccCheckParameter("teamcity_project.bar", "env.MUH", types.ParameterSpec{
// 						Type: types.PasswordType{},
// 					}),
// 					testAccCheckParameter("teamcity_project.bar", "env.TEST", types.ParameterSpec{
// 						Label:       "Test framework",
// 						Description: "Name of the test framework to use",
// 						Type:        types.TextType{"not_empty"},
// 					}),
// 					resource.TestCheckResourceAttr(
// 						"teamcity_project.bar", "parameter.#", "2"),
// 				),
// 			},
// 		},
// 	})
// }

// var testAccProjectParentParameter = `
// resource "teamcity_project" "parent" {
//   name = "Parent"
//   parameter {
//     name = "env.CLOVER"
//     type = "text"
//     validation_mode = "any"
//   }
//   parameter {
//     name = "env.GROVER"
//     type = "text"
//     validation_mode = "any"
//   }
//   parameter_values = {
//     "env.OVER" = "Parent"
//   }
// }
// resource "teamcity_project" "bar" {
//   parent = "${teamcity_project.parent.id}"
//   name = "Bar"
//   parameter {
//     name = "env.OVER"
//     type = "checkbox"
//     checked_value = "Hello"
//   }
//   parameter {
//     name = "env.PLOVER"
//     type = "checkbox"
//     checked_value = "Hello"
//   }
//   parameter_values {
//     "env.OVER" = "Owner"
//   }
// }`

func testAccCheckVcsRpptDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*teamcity.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "teamcity_vcs_root" {
			continue
		}

		// Try to find the Group
		var err error
		vcs, err := client.GetVcsRoot(rs.Primary.ID)

		if err == nil && vcs == nil {
			continue
		}

		if err == nil {
			return fmt.Errorf("VCS Root still exists")
		}

		return err
	}

	return nil
}

func testAccCheckVcsRootExists(n string, v *types.VcsRoot) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Project ID is set")
		}

		client := testAccProvider.Meta().(*teamcity.Client)

		vcs, err := client.GetVcsRoot(rs.Primary.ID)
		if err != nil {
			return err
		}

		if vcs == nil {
			return fmt.Errorf("VCS Root not found")
		}

		*v = *vcs

		return nil
	}
}
