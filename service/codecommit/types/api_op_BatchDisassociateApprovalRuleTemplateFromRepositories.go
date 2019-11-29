// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchDisassociateApprovalRuleTemplateFromRepositoriesInput struct {
	_ struct{} `type:"structure"`

	// The name of the template that you want to disassociate from one or more repositories.
	//
	// ApprovalRuleTemplateName is a required field
	ApprovalRuleTemplateName *string `locationName:"approvalRuleTemplateName" min:"1" type:"string" required:"true"`

	// The repository names that you want to disassociate from the approval rule
	// template.
	//
	// The length constraint limit is for each string in the array. The array itself
	// can be empty.
	//
	// RepositoryNames is a required field
	RepositoryNames []string `locationName:"repositoryNames" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDisassociateApprovalRuleTemplateFromRepositoriesInput"}

	if s.ApprovalRuleTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleTemplateName"))
	}
	if s.ApprovalRuleTemplateName != nil && len(*s.ApprovalRuleTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApprovalRuleTemplateName", 1))
	}

	if s.RepositoryNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput struct {
	_ struct{} `type:"structure"`

	// A list of repository names that have had their association with the template
	// removed.
	//
	// DisassociatedRepositoryNames is a required field
	DisassociatedRepositoryNames []string `locationName:"disassociatedRepositoryNames" type:"list" required:"true"`

	// A list of any errors that might have occurred while attempting to remove
	// the association between the template and the repositories.
	//
	// Errors is a required field
	Errors []BatchDisassociateApprovalRuleTemplateFromRepositoriesError `locationName:"errors" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput) String() string {
	return awsutil.Prettify(s)
}