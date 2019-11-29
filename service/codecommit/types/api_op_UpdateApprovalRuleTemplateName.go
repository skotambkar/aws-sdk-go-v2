// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateApprovalRuleTemplateNameInput struct {
	_ struct{} `type:"structure"`

	// The new name you want to apply to the approval rule template.
	//
	// NewApprovalRuleTemplateName is a required field
	NewApprovalRuleTemplateName *string `locationName:"newApprovalRuleTemplateName" min:"1" type:"string" required:"true"`

	// The current name of the approval rule template.
	//
	// OldApprovalRuleTemplateName is a required field
	OldApprovalRuleTemplateName *string `locationName:"oldApprovalRuleTemplateName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateApprovalRuleTemplateNameInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApprovalRuleTemplateNameInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApprovalRuleTemplateNameInput"}

	if s.NewApprovalRuleTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewApprovalRuleTemplateName"))
	}
	if s.NewApprovalRuleTemplateName != nil && len(*s.NewApprovalRuleTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NewApprovalRuleTemplateName", 1))
	}

	if s.OldApprovalRuleTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OldApprovalRuleTemplateName"))
	}
	if s.OldApprovalRuleTemplateName != nil && len(*s.OldApprovalRuleTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OldApprovalRuleTemplateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateApprovalRuleTemplateNameOutput struct {
	_ struct{} `type:"structure"`

	// The structure and content of the updated approval rule template.
	//
	// ApprovalRuleTemplate is a required field
	ApprovalRuleTemplate *ApprovalRuleTemplate `locationName:"approvalRuleTemplate" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateApprovalRuleTemplateNameOutput) String() string {
	return awsutil.Prettify(s)
}