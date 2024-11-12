package dynatrace_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2023-04-27/tagrules"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type TagRulesResource struct{}

func TestAccDynatraceTagRules_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_dynatrace_tag_rules", "test")
	r := TagRulesResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccDynatraceTagRules_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_dynatrace_tag_rules", "test")
	r := TagRulesResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func (r TagRulesResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := tagrules.ParseTagRuleID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.Dynatrace.TagRulesClient.Get(ctx, *id)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", *id, err)
	}
	return utils.Bool(true), nil
}

func (r TagRulesResource) basic(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
%[1]s

resource "azurerm_dynatrace_tag_rules" "test" {
  name       = "default"
  monitor_id = azurerm_dynatrace_monitor.test.id

  log_rule {
    filtering_tag {
      name   = "Environment"
      value  = "Prod"
      action = "Include"
    }
    send_aad_logs      = "Enabled"
    send_activity_logs = "Enabled"
  }

  metric_rule {
    filtering_tag {
      name   = "Environment"
      value  = "Prod"
      action = "Include"
    }
  }
}
`, template, data.RandomString)
}

func (r TagRulesResource) requiresImport(data acceptance.TestData) string {
	template := r.basic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_dynatrace_tag_rules" "import" {
  name       = azurerm_dynatrace_tag_rules.test.name
  monitor_id = azurerm_dynatrace_tag_rules.test.monitor_id
}
`, template)
}

func (r TagRulesResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%[1]d"
  location = "%[2]s"
}

resource "azurerm_dynatrace_monitor" "test" {
  name                = "acctestacc%[2]s"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  identity {
    type = "SystemAssigned"
  }
  marketplace_subscription = "Active"

  user {
    first_name   = "%s"
    last_name    = "%s"
    email        = "%s"
    phone_number = "%s"
    country      = "%s"
  }

  plan {
    usage_type    = "COMMITTED"
    billing_cycle = "MONTHLY"
    plan          = "azureportalintegration_privatepreview@TIDgmz7xq9ge3py"
  }

  tags = {
    environment = "Prod"
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString, MonitorsResource{}.dynatraceInfo.UserFirstName, MonitorsResource{}.dynatraceInfo.UserLastName, MonitorsResource.dynatraceInfo.UserEmail, MonitorsResource{}.dynatraceInfo.UserPhoneNumber, MonitorsResource{}.dynatraceInfo.UserCountry)
}
