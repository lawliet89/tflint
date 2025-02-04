// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchEventRuleInvalidDescriptionRule checks the pattern is valid
type AwsCloudwatchEventRuleInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsCloudwatchEventRuleInvalidDescriptionRule returns new rule with default attributes
func NewAwsCloudwatchEventRuleInvalidDescriptionRule() *AwsCloudwatchEventRuleInvalidDescriptionRule {
	return &AwsCloudwatchEventRuleInvalidDescriptionRule{
		resourceType:  "aws_cloudwatch_event_rule",
		attributeName: "description",
		max:           512,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventRuleInvalidDescriptionRule) Name() string {
	return "aws_cloudwatch_event_rule_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventRuleInvalidDescriptionRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCloudwatchEventRuleInvalidDescriptionRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventRuleInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventRuleInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
