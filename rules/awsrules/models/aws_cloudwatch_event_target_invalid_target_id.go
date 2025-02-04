// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchEventTargetInvalidTargetIDRule checks the pattern is valid
type AwsCloudwatchEventTargetInvalidTargetIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchEventTargetInvalidTargetIDRule returns new rule with default attributes
func NewAwsCloudwatchEventTargetInvalidTargetIDRule() *AwsCloudwatchEventTargetInvalidTargetIDRule {
	return &AwsCloudwatchEventTargetInvalidTargetIDRule{
		resourceType:  "aws_cloudwatch_event_target",
		attributeName: "target_id",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^[\.\-_A-Za-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventTargetInvalidTargetIDRule) Name() string {
	return "aws_cloudwatch_event_target_invalid_target_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventTargetInvalidTargetIDRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCloudwatchEventTargetInvalidTargetIDRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventTargetInvalidTargetIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventTargetInvalidTargetIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"target_id must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"target_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`target_id does not match valid pattern ^[\.\-_A-Za-z0-9]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
