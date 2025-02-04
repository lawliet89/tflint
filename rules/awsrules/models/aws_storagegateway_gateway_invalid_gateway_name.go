// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsStoragegatewayGatewayInvalidGatewayNameRule checks the pattern is valid
type AwsStoragegatewayGatewayInvalidGatewayNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsStoragegatewayGatewayInvalidGatewayNameRule returns new rule with default attributes
func NewAwsStoragegatewayGatewayInvalidGatewayNameRule() *AwsStoragegatewayGatewayInvalidGatewayNameRule {
	return &AwsStoragegatewayGatewayInvalidGatewayNameRule{
		resourceType:  "aws_storagegateway_gateway",
		attributeName: "gateway_name",
		max:           255,
		min:           2,
		pattern:       regexp.MustCompile(`^[ -\.0-\[\]-~]*[!-\.0-\[\]-~][ -\.0-\[\]-~]*$`),
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayGatewayInvalidGatewayNameRule) Name() string {
	return "aws_storagegateway_gateway_invalid_gateway_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayGatewayInvalidGatewayNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsStoragegatewayGatewayInvalidGatewayNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayGatewayInvalidGatewayNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayGatewayInvalidGatewayNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"gateway_name must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"gateway_name must be 2 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`gateway_name does not match valid pattern ^[ -\.0-\[\]-~]*[!-\.0-\[\]-~][ -\.0-\[\]-~]*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
