// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchEventTargetInvalidInputPathRule checks the pattern is valid
type AwsCloudwatchEventTargetInvalidInputPathRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsCloudwatchEventTargetInvalidInputPathRule returns new rule with default attributes
func NewAwsCloudwatchEventTargetInvalidInputPathRule() *AwsCloudwatchEventTargetInvalidInputPathRule {
	return &AwsCloudwatchEventTargetInvalidInputPathRule{
		resourceType:  "aws_cloudwatch_event_target",
		attributeName: "input_path",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventTargetInvalidInputPathRule) Name() string {
	return "aws_cloudwatch_event_target_invalid_input_path"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventTargetInvalidInputPathRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCloudwatchEventTargetInvalidInputPathRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventTargetInvalidInputPathRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventTargetInvalidInputPathRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"input_path must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
