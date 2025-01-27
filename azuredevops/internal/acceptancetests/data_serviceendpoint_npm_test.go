//go:build (all || data_sources || data_serviceendpoint_npm) && (!exclude_data_sources || !exclude_data_serviceendpoint_npm)
// +build all data_sources data_serviceendpoint_npm
// +build !exclude_data_sources !exclude_data_serviceendpoint_npm

package acceptancetests

import (
	"fmt"
	"testing"

	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccServiceEndpointNpm_DataSource(t *testing.T) {
	name := testutils.GenerateResourceName()

	tfNode := "data.azuredevops_serviceendpoint_npm.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutils.PreCheck(t, nil) },
		Providers: testutils.GetProviders(),
		Steps: []resource.TestStep{
			{
				Config: hclServiceEndpointNpmDataSource(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tfNode, "service_endpoint_name", name),
				),
			},
		},
	})
}

func hclServiceEndpointNpmDataSource(name string) string {
	return fmt.Sprintf(`
resource "azuredevops_project" "test" {
  name               = "%[1]s"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
}

resource "azuredevops_serviceendpoint_npm" "test" {
  project_id            = azuredevops_project.test.id
  service_endpoint_name = "%[1]s"
  access_token          = "redacted"
  url                   = "http://url.com/"
}

data "azuredevops_serviceendpoint_npm" "test" {
  project_id          = azuredevops_project.test.id
  service_endpoint_name = azuredevops_serviceendpoint_npm.test.service_endpoint_name
}
`, name)
}
