// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteOrganizationConfigRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of organization config rule that you want to delete.
	//
	// OrganizationConfigRuleName is a required field
	OrganizationConfigRuleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteOrganizationConfigRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteOrganizationConfigRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteOrganizationConfigRuleInput"}

	if s.OrganizationConfigRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationConfigRuleName"))
	}
	if s.OrganizationConfigRuleName != nil && len(*s.OrganizationConfigRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationConfigRuleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteOrganizationConfigRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOrganizationConfigRuleOutput) String() string {
	return awsutil.Prettify(s)
}