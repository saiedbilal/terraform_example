// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynatrace

import "github.com/hashicorp/terraform-provider-azurerm/internal/sdk"

type Registration struct{}

var _ sdk.TypedServiceRegistrationWithAGitHubLabel = Registration{}

func (r Registration) AssociatedGitHubLabel() string {
	return "service/dynatrace"
}

func (r Registration) Name() string {
	return "Dynatrace"
}

func (r Registration) DataSources() []sdk.DataSource {
	return []sdk.DataSource{}
}

func (r Registration) Resources() []sdk.Resource {
	return []sdk.Resource{
		MonitorsResource{},
<<<<<<< HEAD
=======
		TagRulesResource{},
>>>>>>> 1f5e5db9b0 (New resources: azurerm_dynatrace_monitors, azurerm_dynatrace_tag_rules)
	}
}

func (r Registration) WebsiteCategories() []string {
	return []string{
		"Dynatrace",
	}
}
