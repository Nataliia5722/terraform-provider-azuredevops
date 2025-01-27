//go:build all || core
// +build all core

package acceptancetests

import (
	"os"
	"testing"

	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Verifies that the client config data source loads the configured AzDO org
func TestAccClientConfig_LoadsCorrectProperties(t *testing.T) {
	tfNode := "data.azuredevops_client_config.c"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutils.PreCheck(t, nil) },
		Providers: testutils.GetProviders(),
		Steps: []resource.TestStep{
			{
				Config: `data "azuredevops_client_config" "c" {}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(tfNode, "id"),
					resource.TestCheckResourceAttr(tfNode, "organization_url", os.Getenv("AZDO_ORG_SERVICE_URL")),
				),
			},
		},
	})
}
