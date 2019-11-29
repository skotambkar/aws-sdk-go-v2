// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteApprovalRuleTemplateInput struct {
	_ struct{} `type:"structure"`

	// The name of the approval rule template to delete.
	//
	// ApprovalRuleTemplateName is a required field
	ApprovalRuleTemplateName *string `locationName:"approvalRuleTemplateName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApprovalRuleTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApprovalRuleTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApprovalRuleTemplateInput"}

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

type DeleteApprovalRuleTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The system-generated ID of the deleted approval rule template. If the template
	// has been previously deleted, the only response is a 200 OK.
	//
	// ApprovalRuleTemplateId is a required field
	ApprovalRuleTemplateId *string `locationName:"approvalRuleTemplateId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApprovalRuleTemplateOutput) String() string {
	return awsutil.Prettify(s)
}
