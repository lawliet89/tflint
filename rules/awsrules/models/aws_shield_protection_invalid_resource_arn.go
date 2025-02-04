// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsShieldProtectionInvalidResourceArnRule checks the pattern is valid
type AwsShieldProtectionInvalidResourceArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsShieldProtectionInvalidResourceArnRule returns new rule with default attributes
func NewAwsShieldProtectionInvalidResourceArnRule() *AwsShieldProtectionInvalidResourceArnRule {
	return &AwsShieldProtectionInvalidResourceArnRule{
		resourceType:  "aws_shield_protection",
		attributeName: "resource_arn",
		max:           2048,
		min:           1,
		pattern:       regexp.MustCompile(`^arn:aws.*`),
	}
}

// Name returns the rule name
func (r *AwsShieldProtectionInvalidResourceArnRule) Name() string {
	return "aws_shield_protection_invalid_resource_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsShieldProtectionInvalidResourceArnRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsShieldProtectionInvalidResourceArnRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsShieldProtectionInvalidResourceArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsShieldProtectionInvalidResourceArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"resource_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"resource_arn must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`resource_arn does not match valid pattern ^arn:aws.*`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
