// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynatrace

import (
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2023-04-27/monitors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2023-04-27/tagrules"
)

func ExpandDynatracePlanData(input []PlanData) *monitors.PlanData {
	if len(input) == 0 {
		return nil
	}
	v := input[0]

	return pointer.To(monitors.PlanData{
		BillingCycle: &v.BillingCycle,
		PlanDetails:  &v.PlanDetails,
		UsageType:    &v.UsageType,
	})
}

func ExpandDynatraceUserInfo(input []UserInfo) *monitors.UserInfo {
	if len(input) == 0 {
		return nil
	}
	v := input[0]

	return pointer.To(monitors.UserInfo{
		Country:      pointer.To(v.Country),
		EmailAddress: pointer.To(v.EmailAddress),
		FirstName:    pointer.To(v.FirstName),
		LastName:     pointer.To(v.LastName),
		PhoneNumber:  pointer.To(v.PhoneNumber),
	})
}

func FlattenDynatracePlanData(input *monitors.PlanData) []PlanData {
	if input == nil {
		return []PlanData{}
	}

	var billingCycle string
	var effectiveDate string
	var planDetails string
	var usageType string

	if input.BillingCycle != nil {
		billingCycle = pointer.From(input.BillingCycle)
	}

	if input.EffectiveDate != nil {
		effectiveDate = pointer.From(input.EffectiveDate)
	}

	if input.PlanDetails != nil {
		planDetails = pointer.From(input.PlanDetails)
	}

	if input.UsageType != nil {
		usageType = pointer.From(input.UsageType)
	}

	return []PlanData{
		{
			BillingCycle:  billingCycle,
			EffectiveDate: effectiveDate,
			PlanDetails:   planDetails,
			UsageType:     usageType,
		},
	}
}

func FlattenDynatraceUserInfo(input []interface{}) []UserInfo {
	if len(input) == 0 {
		return []UserInfo{}
	}

	v := input[0].(map[string]interface{})
	return []UserInfo{
		{
			Country:      v["country"].(string),
			EmailAddress: v["email"].(string),
			FirstName:    v["first_name"].(string),
			LastName:     v["last_name"].(string),
			PhoneNumber:  v["phone_number"].(string),
		},
	}
}

func FlattenLogRules(input *tagrules.LogRules) []LogRule {
	if input == nil {
		return []LogRule{}
	}

	var filteringTags []FilteringTag
	var sendAadLogs string
	var sendActivityLogs string
	var sendSubscriptionLogs string

	if input.FilteringTags != nil {
		filteringTags = FlattenFilteringTags(input.FilteringTags)
	}

	if input.SendActivityLogs != nil {
		sendActivityLogs = string(*input.SendActivityLogs)
	}

	if input.SendAadLogs != nil {
		sendAadLogs = string(*input.SendAadLogs)
	}

	if input.SendSubscriptionLogs != nil {
		sendSubscriptionLogs = string(*input.SendSubscriptionLogs)
	}

	return []LogRule{
		{
			FilteringTags:        filteringTags,
			SendAadLogs:          sendAadLogs,
			SendActivityLogs:     sendActivityLogs,
			SendSubscriptionLogs: sendSubscriptionLogs,
		},
	}
}

func FlattenFilteringTags(input *[]tagrules.FilteringTag) []FilteringTag {
	if input == nil || len(*input) == 0 {
		return []FilteringTag{}
	}

	var name string
	var value string
	var action string
	tags := *input
	v := tags[0]

	if v.Name != nil {
		name = *v.Name
	}

	if v.Value != nil {
		value = *v.Value
	}

	if v.Action != nil {
		action = string(*v.Action)
	}

	return []FilteringTag{
		{
			Name:   name,
			Value:  value,
			Action: action,
		},
	}
}

func FlattenMetricRules(input *tagrules.MetricRules) []MetricRule {
	if input == nil {
		return []MetricRule{}
	}

	var filteringTags []FilteringTag

	if input.FilteringTags != nil {
		filteringTags = FlattenFilteringTags(input.FilteringTags)
	}

	return []MetricRule{
		{
			FilteringTags: filteringTags,
		},
	}
}

func ExpandMetricRules(input []MetricRule) *tagrules.MetricRules {
	if len(input) == 0 {
		return nil
	}
	v := input[0]

	return &tagrules.MetricRules{
		FilteringTags: ExpandFilteringTag(v.FilteringTags),
	}
}

func ExpandLogRule(input []LogRule) *tagrules.LogRules {
	if len(input) == 0 {
		return nil
	}
	v := input[0]
	sendAadLogs := tagrules.SendAadLogsStatus(v.SendAadLogs)
	sendActivityLogs := tagrules.SendActivityLogsStatus(v.SendActivityLogs)
	sendSubscriptionLogs := tagrules.SendSubscriptionLogsStatus(v.SendSubscriptionLogs)

	return &tagrules.LogRules{
		FilteringTags:        ExpandFilteringTag(v.FilteringTags),
		SendAadLogs:          pointer.To(sendAadLogs),
		SendActivityLogs:     pointer.To(sendActivityLogs),
		SendSubscriptionLogs: pointer.To(sendSubscriptionLogs),
	}
}

func ExpandFilteringTag(input []FilteringTag) *[]tagrules.FilteringTag {
	if len(input) == 0 {
		return nil
	}
	v := input[0]
	action := tagrules.TagAction(v.Action)

	return &[]tagrules.FilteringTag{
		{
			Action: pointer.To(action),
			Name:   pointer.To(v.Name),
			Value:  pointer.To(v.Value),
		},
	}
}
