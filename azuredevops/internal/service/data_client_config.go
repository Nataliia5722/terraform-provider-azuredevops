package service

import (
	"time"

	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	organizationURL = "organization_url"
)

// DataClientConfig schema and implementation for AzDO client configuration
func DataClientConfig() *schema.Resource {
	return &schema.Resource{
		Read: clientConfigRead,
		Schema: map[string]*schema.Schema{
			organizationURL: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func clientConfigRead(d *schema.ResourceData, m interface{}) error {
	// The ID is meaningless for this data source, so ID can act as a
	// point in time snapshot
	d.SetId(time.Now().UTC().String())
	d.Set(organizationURL, m.(*client.AggregatedClient).OrganizationURL)
	return nil
}
