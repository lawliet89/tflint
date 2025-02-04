// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsOpsworksStackInvalidDefaultRootDeviceTypeRule checks the pattern is valid
type AwsOpsworksStackInvalidDefaultRootDeviceTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsOpsworksStackInvalidDefaultRootDeviceTypeRule returns new rule with default attributes
func NewAwsOpsworksStackInvalidDefaultRootDeviceTypeRule() *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule {
	return &AwsOpsworksStackInvalidDefaultRootDeviceTypeRule{
		resourceType:  "aws_opsworks_stack",
		attributeName: "default_root_device_type",
		enum: []string{
			"ebs",
			"instance-store",
		},
	}
}

// Name returns the rule name
func (r *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule) Name() string {
	return "aws_opsworks_stack_invalid_default_root_device_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule) Check(runner *tflint.Runner) error {
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
					`default_root_device_type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
