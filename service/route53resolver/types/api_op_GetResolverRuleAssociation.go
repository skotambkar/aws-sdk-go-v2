// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetResolverRuleAssociationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the resolver rule association that you want to get information
	// about.
	//
	// ResolverRuleAssociationId is a required field
	ResolverRuleAssociationId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetResolverRuleAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResolverRuleAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResolverRuleAssociationInput"}

	if s.ResolverRuleAssociationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverRuleAssociationId"))
	}
	if s.ResolverRuleAssociationId != nil && len(*s.ResolverRuleAssociationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverRuleAssociationId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetResolverRuleAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the resolver rule association that you specified in a GetResolverRuleAssociation
	// request.
	ResolverRuleAssociation *ResolverRuleAssociation `type:"structure"`
}

// String returns the string representation
func (s GetResolverRuleAssociationOutput) String() string {
	return awsutil.Prettify(s)
}