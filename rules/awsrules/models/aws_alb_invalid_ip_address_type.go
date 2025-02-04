// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsALBInvalidIPAddressTypeRule checks the pattern is valid
type AwsALBInvalidIPAddressTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsALBInvalidIPAddressTypeRule returns new rule with default attributes
func NewAwsALBInvalidIPAddressTypeRule() *AwsALBInvalidIPAddressTypeRule {
	return &AwsALBInvalidIPAddressTypeRule{
		resourceType:  "aws_alb",
		attributeName: "ip_address_type",
		enum: []string{
			"ipv4",
			"dualstack",
		},
	}
}

// Name returns the rule name
func (r *AwsALBInvalidIPAddressTypeRule) Name() string {
	return "aws_alb_invalid_ip_address_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsALBInvalidIPAddressTypeRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsALBInvalidIPAddressTypeRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsALBInvalidIPAddressTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsALBInvalidIPAddressTypeRule) Check(runner *tflint.Runner) error {
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
					`ip_address_type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
