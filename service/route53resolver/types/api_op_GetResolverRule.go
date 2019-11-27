// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetResolverRuleInput struct {
	_ struct{} `type:"structure"`

	// The ID of the resolver rule that you want to get information about.
	//
	// ResolverRuleId is a required field
	ResolverRuleId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetResolverRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResolverRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResolverRuleInput"}

	if s.ResolverRuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverRuleId"))
	}
	if s.ResolverRuleId != nil && len(*s.ResolverRuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverRuleId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetResolverRuleOutput struct {
	_ struct{} `type:"structure"`

	// Information about the resolver rule that you specified in a GetResolverRule
	// request.
	ResolverRule *ResolverRule `type:"structure"`
}

// String returns the string representation
func (s GetResolverRuleOutput) String() string {
	return awsutil.Prettify(s)
}
