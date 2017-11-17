package teamcity

import (
	// "fmt"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/Cardfree/teamcity-sdk-go/teamcity"
	"github.com/Cardfree/teamcity-sdk-go/types"
	"log"
)

func resourceAgentPoolAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAgentPoolAttachementCreate,
		Read:   resourceAgentPoolAttachementRead,
		Delete: resourceAgentPoolAttachementDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAgentPoolAttachementCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*teamcity.Client)

	attachmentPool := d.Get("pool").(string)
	attachmentProject := d.Get("project").(string)

	attachment := types.AgentPoolAttachment{
		ProjectID: attachmentProject,
	}
	err := client.CreateAgentPoolAttachment(attachmentPool, &attachment)
	if err != nil {
		return err
	}

	// THIS ID NEEDS TO BE CHANGED TO SOMETHING COMBINING POOL AND PROJECT
	// id := attachment.ID
	d.SetId(attachmentProject)
	return nil
}

func resourceAgentPoolAttachementRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Reading agent_attachment resource %q", d.Id())
	client := meta.(*teamcity.Client)
	pool, err := client.GetAgentPool(d.Get("pool").(string))
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Read Agent Pools: %q", pool)
	var project types.Project
	project = pool.Projects[d.Get("project").(string)]

	log.Printf("[DEBUG] Read Agent Pool: %q", project)

	if project.ID == "" {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceAgentPoolAttachementDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*teamcity.Client)
	return client.DeleteAgentPoolAttachement(d.Get("name").(string), d.Get("pool").(string))
	// return nil
}
