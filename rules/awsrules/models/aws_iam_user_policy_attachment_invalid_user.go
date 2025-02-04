// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsIAMUserPolicyAttachmentInvalidUserRule checks the pattern is valid
type AwsIAMUserPolicyAttachmentInvalidUserRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMUserPolicyAttachmentInvalidUserRule returns new rule with default attributes
func NewAwsIAMUserPolicyAttachmentInvalidUserRule() *AwsIAMUserPolicyAttachmentInvalidUserRule {
	return &AwsIAMUserPolicyAttachmentInvalidUserRule{
		resourceType:  "aws_iam_user_policy_attachment",
		attributeName: "user",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMUserPolicyAttachmentInvalidUserRule) Name() string {
	return "aws_iam_user_policy_attachment_invalid_user"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMUserPolicyAttachmentInvalidUserRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsIAMUserPolicyAttachmentInvalidUserRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMUserPolicyAttachmentInvalidUserRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMUserPolicyAttachmentInvalidUserRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"user must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"user must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`user does not match valid pattern ^[\w+=,.@-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
