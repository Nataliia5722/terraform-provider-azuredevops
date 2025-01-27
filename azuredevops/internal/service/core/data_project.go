package core

import (
	"context"
	"fmt"

	"github.com/Nataliia5722/azure-devops-go-api/azuredevops/v7/core"
	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/client"
	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/utils"
	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/utils/converter"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// DataProject schema and implementation for project data source
func DataProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataProjectRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"project_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
				ConflictsWith: []string{
					"name",
				},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"visibility": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version_control": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"work_item_template": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"process_template_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"features": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

// Introducing a read method here which is almost the same code a in resource_project.go
// but this follows the `A little copying is better than a little dependency.` GO proverb.
func dataProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clients := m.(*client.AggregatedClient)

	name := d.Get("name").(string)
	id := d.Get("project_id").(string)

	if name == "" && id == "" {
		return diag.FromErr(fmt.Errorf("Either project_id or name must be set "))
	}

	identifier := id
	if identifier == "" {
		identifier = name
	}

	project, err := clients.CoreClient.GetProject(ctx, core.GetProjectArgs{
		ProjectId:           &identifier,
		IncludeCapabilities: converter.Bool(true),
		IncludeHistory:      converter.Bool(false),
	})

	if err != nil {
		if utils.ResponseWasNotFound(err) {
			return diag.FromErr(fmt.Errorf("Project with name %s or ID %s does not exist ", name, id))
		}
		return diag.FromErr(fmt.Errorf("Error looking up project with Name %s or ID %s, %+v ", name, id, err))
	}

	err = flattenProject(clients, d, project)
	d.Set("project_id", project.Id.String())
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error flattening project: %v", err))
	}
	return nil
}
