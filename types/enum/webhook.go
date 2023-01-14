// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package enum

// WebhookParent defines different types of parents of a webhook.
type WebhookParent string

func (WebhookParent) Enum() []interface{} { return toInterfaceSlice(webhookParents) }

const (
	// WebhookParentRepo describes a repo as webhook owner.
	WebhookParentRepo WebhookParent = "repo"

	// WebhookParentSpace describes a space as webhook owner.
	WebhookParentSpace WebhookParent = "space"
)

var webhookParents = sortEnum([]WebhookParent{
	WebhookParentRepo,
	WebhookParentSpace,
})

// WebhookExecutionResult defines the different results of a webhook execution.
type WebhookExecutionResult string

func (WebhookExecutionResult) Enum() []interface{} { return toInterfaceSlice(webhookExecutionResults) }

const (
	// WebhookExecutionResultSuccess describes a webhook execution result that succeeded.
	WebhookExecutionResultSuccess WebhookExecutionResult = "success"

	// WebhookExecutionResultRetriableError describes a webhook execution result that failed with a retriable error.
	WebhookExecutionResultRetriableError WebhookExecutionResult = "retriable_error"

	// WebhookExecutionResultFatalError describes a webhook execution result that failed with an unrecoverable error.
	WebhookExecutionResultFatalError WebhookExecutionResult = "fatal_error"
)

var webhookExecutionResults = sortEnum([]WebhookExecutionResult{
	WebhookExecutionResultSuccess,
	WebhookExecutionResultRetriableError,
	WebhookExecutionResultFatalError,
})

// WebhookTrigger defines the different types of webhook triggers available.
type WebhookTrigger string

func (WebhookTrigger) Enum() []interface{}                { return toInterfaceSlice(webhookTriggers) }
func (s WebhookTrigger) Sanitize() (WebhookTrigger, bool) { return Sanitize(s, GetAllWebhookTriggers) }

func GetAllWebhookTriggers() ([]WebhookTrigger, WebhookTrigger) {
	return webhookTriggers, "" // No default value
}

const (
	// WebhookTriggerBranchCreated gets triggered when a branch gets created.
	WebhookTriggerBranchCreated WebhookTrigger = "branch_created"
	// WebhookTriggerBranchUpdated gets triggered when a branch gets updated.
	WebhookTriggerBranchUpdated WebhookTrigger = "branch_updated"
	// WebhookTriggerBranchDeleted gets triggered when a branch gets deleted.
	WebhookTriggerBranchDeleted WebhookTrigger = "branch_deleted"

	// WebhookTriggerTagCreated gets triggered when a tag gets created.
	WebhookTriggerTagCreated WebhookTrigger = "tag_created"
	// WebhookTriggerTagUpdated gets triggered when a tag gets updated.
	WebhookTriggerTagUpdated WebhookTrigger = "tag_updated"
	// WebhookTriggerTagDeleted gets triggered when a tag gets deleted.
	WebhookTriggerTagDeleted WebhookTrigger = "tag_deleted"
)

var webhookTriggers = sortEnum([]WebhookTrigger{
	WebhookTriggerBranchCreated,
	WebhookTriggerBranchUpdated,
	WebhookTriggerBranchDeleted,
	WebhookTriggerTagCreated,
	WebhookTriggerTagUpdated,
	WebhookTriggerTagDeleted,
})