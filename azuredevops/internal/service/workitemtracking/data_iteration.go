package workitemtracking

import (
	"github.com/Nataliia5722/azure-devops-go-api/azuredevops/v7/workitemtracking"
	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/client"
	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/service/workitemtracking/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataIteration schema for iteration data
func DataIteration() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceIterationRead,
		Schema: utils.CreateClassificationNodeSchema(map[string]*schema.Schema{}),
	}
}

func dataSourceIterationRead(d *schema.ResourceData, m interface{}) error {
	clients := m.(*client.AggregatedClient)
	return utils.ReadClassificationNode(clients, d, workitemtracking.TreeStructureGroupValues.Iterations)
}
