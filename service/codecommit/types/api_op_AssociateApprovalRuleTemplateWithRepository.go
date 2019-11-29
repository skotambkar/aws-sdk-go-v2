// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateApprovalRuleTemplateWithRepositoryInput struct {
	_ struct{} `type:"structure"`

	// The name for the approval rule template.
	//
	// ApprovalRuleTemplateName is a required field
	ApprovalRuleTemplateName *string `locationName:"approvalRuleTemplateName" min:"1" type:"string" required:"true"`

	// The name of the repository that you want to associate with the template.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateApprovalRuleTemplateWithRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateApprovalRuleTemplateWithRepositoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateApprovalRuleTemplateWithRepositoryInput"}

	if s.ApprovalRuleTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleTemplateName"))
	}
	if s.ApprovalRuleTemplateName != nil && len(*s.ApprovalRuleTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApprovalRuleTemplateName", 1))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateApprovalRuleTemplateWithRepositoryOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateApprovalRuleTemplateWithRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}
