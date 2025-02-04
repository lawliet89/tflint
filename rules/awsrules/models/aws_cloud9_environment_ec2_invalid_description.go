// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCloud9EnvironmentEc2InvalidDescriptionRule checks the pattern is valid
type AwsCloud9EnvironmentEc2InvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsCloud9EnvironmentEc2InvalidDescriptionRule returns new rule with default attributes
func NewAwsCloud9EnvironmentEc2InvalidDescriptionRule() *AwsCloud9EnvironmentEc2InvalidDescriptionRule {
	return &AwsCloud9EnvironmentEc2InvalidDescriptionRule{
		resourceType:  "aws_cloud9_environment_ec2",
		attributeName: "description",
		max:           200,
	}
}

// Name returns the rule name
func (r *AwsCloud9EnvironmentEc2InvalidDescriptionRule) Name() string {
	return "aws_cloud9_environment_ec2_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloud9EnvironmentEc2InvalidDescriptionRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCloud9EnvironmentEc2InvalidDescriptionRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCloud9EnvironmentEc2InvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloud9EnvironmentEc2InvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 200 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
