// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateApprovalRuleTemplateDescriptionInput struct {
	_ struct{} `type:"structure"`

	// The updated description of the approval rule template.
	//
	// ApprovalRuleTemplateDescription is a required field
	ApprovalRuleTemplateDescription *string `locationName:"approvalRuleTemplateDescription" type:"string" required:"true"`

	// The name of the template for which you want to update the description.
	//
	// ApprovalRuleTemplateName is a required field
	ApprovalRuleTemplateName *string `locationName:"approvalRuleTemplateName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateApprovalRuleTemplateDescriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApprovalRuleTemplateDescriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApprovalRuleTemplateDescriptionInput"}

	if s.ApprovalRuleTemplateDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleTemplateDescription"))
	}

	if s.ApprovalRuleTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleTemplateName"))
	}
	if s.ApprovalRuleTemplateName != nil && len(*s.ApprovalRuleTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApprovalRuleTemplateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateApprovalRuleTemplateDescriptionOutput struct {
	_ struct{} `type:"structure"`

	// The structure and content of the updated approval rule template.
	//
	// ApprovalRuleTemplate is a required field
	ApprovalRuleTemplate *ApprovalRuleTemplate `locationName:"approvalRuleTemplate" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateApprovalRuleTemplateDescriptionOutput) String() string {
	return awsutil.Prettify(s)
}