//go:build (all || core || resource_project || resource_project_features) && !exclude_resource_project_features
// +build all core resource_project resource_project_features
// +build !exclude_resource_project_features

package acceptancetests

import (
	"testing"

	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccProjectFeatures_EnableUpdateFeature(t *testing.T) {
	projectName := testutils.GenerateResourceName()
	tfNode := "azuredevops_project_features.project-features"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testutils.HclProjectFeatures(projectName, "disabled", "disabled"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tfNode, "features.testplans", "disabled"),
					resource.TestCheckResourceAttr(tfNode, "features.artifacts", "disabled"),
				),
				Destroy: false,
			},
			{
				Config: testutils.HclProjectFeatures(projectName, "enabled", "disabled"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tfNode, "features.testplans", "enabled"),
					resource.TestCheckResourceAttr(tfNode, "features.artifacts", "disabled"),
				),
			},
		},
	})
}
