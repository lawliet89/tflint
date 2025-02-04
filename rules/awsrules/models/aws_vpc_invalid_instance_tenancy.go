// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsVpcInvalidInstanceTenancyRule checks the pattern is valid
type AwsVpcInvalidInstanceTenancyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsVpcInvalidInstanceTenancyRule returns new rule with default attributes
func NewAwsVpcInvalidInstanceTenancyRule() *AwsVpcInvalidInstanceTenancyRule {
	return &AwsVpcInvalidInstanceTenancyRule{
		resourceType:  "aws_vpc",
		attributeName: "instance_tenancy",
		enum: []string{
			"default",
			"dedicated",
			"host",
		},
	}
}

// Name returns the rule name
func (r *AwsVpcInvalidInstanceTenancyRule) Name() string {
	return "aws_vpc_invalid_instance_tenancy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsVpcInvalidInstanceTenancyRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsVpcInvalidInstanceTenancyRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsVpcInvalidInstanceTenancyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsVpcInvalidInstanceTenancyRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`instance_tenancy is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
