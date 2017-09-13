package main

import (
	// "fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"errors"
	"github.com/Cardfree/teamcity-sdk-go/teamcity"
	"github.com/Cardfree/teamcity-sdk-go/types"
	"log"
	// "reflect"
)

func resourceVcsRoot() *schema.Resource {
	return &schema.Resource{
		Create: resourceVcsRootCreate,
		Read:   resourceVcsRootRead,
		Update: resourceVcsRootUpdate,
		Delete: resourceVcsRootDelete,

		Schema: map[string]*schema.Schema{
			"project": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: teamcity.ValidateID,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcs_provider": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"properties": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

/*
   ID         string     `json:"id"`
   Name       string     `json:"name"`
   VcsName    string     `json:"vcsName"`
   Href       string     `json:"href"`
   ProjectID  ProjectId  `json:"project"`
  Properties Properties `json:"properties"`
*/

func resourceVcsRootCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*teamcity.Client)

	d.Partial(true)
	project := d.Get("project").(string)
	name := d.Get("name").(string)
	vcs_provider := d.Get("vcs_provider").(string)
	if vcs_provider == "" {
		vcs_provider = "jetbrains.git"
	}
	if project == "" {
		project = "_Root"
	}

	vcs := types.VcsRoot{
		ProjectID: project,
		Name:      name,
		VcsName:   vcs_provider,
	}
	err := client.CreateVcsRoot(&vcs)
	if err != nil {
		return err
	}
	id := vcs.ID
	d.SetId(id)
	d.SetPartial("project")
	d.SetPartial("name")
	d.SetPartial("vcs_provider")

	// if parent == "" {
	// 	parent = "_Root"
	// }
	// var parent_parameters types.Parameters
	// if parent_project, err := client.GetProject(parent); err != nil {
	// 	return err
	// } else {
	// 	parent_parameters = parent_project.Parameters
	// }

	// parameters := definitionToParameters(*d.Get("parameter").(*schema.Set))
	// for name, _ := range parameters {
	// 	if parent_parameter, ok := parent_parameters[name]; ok && parent_parameter.Spec != nil {
	// 		return fmt.Errorf("Can't redefine parent parameter %s", name)
	// 	}
	// }
	// for name, v := range d.Get("parameter_values").(map[string]interface{}) {
	// 	value := v.(string)
	// 	parameter, ok := parameters[name]
	// 	if !ok {
	// 		if parameter, ok = parent_parameters[name]; !ok {
	// 			parameter = types.Parameter{
	// 				Value: value,
	// 			}
	// 		}
	// 	}
	// 	parameter.Value = value
	// 	parameters[name] = parameter
	// 	log.Printf("Parameter value %s => %s", name, parameter.Value)
	// }
	// log.Printf("Replace Parameters value %q", parameters)
	// if err := client.ReplaceAllProjectParameters(id, &parameters); err != nil {
	// 	return err
	// }
	// d.SetPartial("parameter_values")
	// d.SetPartial("parameter")

	d.Partial(false)
	return nil
}

func resourceVcsRootRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Reading vcs_root resource %q", d.Id())
	client := meta.(*teamcity.Client)
	vcs, err := client.GetVcsRoot(d.Id())
	if err != nil {
		return err
	}
	if vcs == nil {
		d.SetId("")
		return nil
	}

	// parent := vcs.ParentProjectID
	// if parent == "_Root" {
	// 	parent = ""
	// }

	d.Set("project", vcs.ProjectID)
	d.Set("name", vcs.Name)
	d.Set("vcs_provider", vcs.VcsName)

	// var parent_parameters types.Parameters
	// if parent_project, err := client.GetProject(string(project.ParentProjectID)); err != nil {
	// 	return err
	// } else {
	// 	parent_parameters = parent_project.Parameters
	// }
	// parameters := project.Parameters
	// values := make(map[string]interface{})
	// current := d.Get("parameter_values").(map[string]interface{})
	// for name, parameter := range project.Parameters {
	// 	if parent_parameter, ok := parent_parameters[name]; ok {
	// 		if parent_parameter.Value != parameter.Value {
	// 			values[name] = parameter.Value
	// 		}
	// 		if parent_parameter.Spec != nil || parameter.Spec == nil {
	// 			delete(parameters, name)
	// 		}
	// 	} else {
	// 		if parameter.Spec == nil {
	// 			delete(parameters, name)
	// 		}
	// 		pwt := types.PasswordType{}
	// 		if parameter.Value != "" {
	// 			values[name] = parameter.Value
	// 		} else if parameter.Spec != nil && parameter.Spec.Type == pwt {
	// 			if value, ok := current[name]; ok && value != "" {
	// 				values[name] = value
	// 			}
	// 		}
	// 	}
	// }
	// d.Set("parameter", parametersToDefinition(parameters))
	// d.Set("parameter_values", values)

	return nil
}

func resourceVcsRootUpdate(d *schema.ResourceData, meta interface{}) error {
	return errors.New("Update is not supported")
	// client := meta.(*teamcity.Client)

	// id := d.Id()
	// d.Partial(true)

	// if d.HasChange("description") {
	// 	if err := client.SetProjectDescription(d.Id(), d.Get("description").(string)); err != nil {
	// 		return err
	// 	}
	// 	d.SetPartial("description")
	// }

	// if d.HasChange("parameter") || d.HasChange("parameter_values") {
	// 	parent := d.Get("parent").(string)
	// 	if parent == "" {
	// 		parent = "_Root"
	// 	}
	// 	var parent_parameters types.Parameters
	// 	if parent_project, err := client.GetProject(parent); err != nil {
	// 		return err
	// 	} else {
	// 		parent_parameters = parent_project.Parameters
	// 	}

	// 	o, n := d.GetChange("parameter")
	// 	parameters := definitionToParameters(*n.(*schema.Set))
	// 	old := definitionToParameters(*o.(*schema.Set))
	// 	replace_parameters := make(types.Parameters)
	// 	delete_parameters := old
	// 	for name, parameter := range parameters {
	// 		if parent_parameter, ok := parent_parameters[name]; ok && parent_parameter.Spec != nil {
	// 			return fmt.Errorf("Can't redefine parent parameter %s", name)
	// 		}
	// 		if !reflect.DeepEqual(parameter, old[name]) {
	// 			replace_parameters[name] = parameter
	// 		}
	// 		delete(delete_parameters, name)
	// 	}
	// 	for name, v := range d.Get("parameter_values").(map[string]interface{}) {
	// 		value := v.(string)
	// 		parameter, ok := parameters[name]
	// 		if !ok {
	// 			if parameter, ok = parent_parameters[name]; !ok {
	// 				parameter = types.Parameter{
	// 					Value: value,
	// 				}
	// 			}
	// 		}
	// 		parameter.Value = value
	// 		replace_parameters[name] = parameter
	// 	}
	// 	for name, _ := range delete_parameters {
	// 		if err := client.DeleteProjectParameter(id, name); err != nil {
	// 			return err
	// 		}
	// 	}
	// 	for name, parameter := range replace_parameters {
	// 		if err := client.ReplaceProjectParameter(id, name, &parameter); err != nil {
	// 			return err
	// 		}
	// 	}
	// 	d.SetPartial("parameter_values")
	// 	d.SetPartial("parameter")
	// }

	return nil
}

func resourceVcsRootDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*teamcity.Client)
	return client.DeleteVcsRoot(d.Id())
}
