// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the DisableTopicRuleRequest operation.
type DisableTopicRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of the rule to disable.
	//
	// RuleName is a required field
	RuleName *string `location:"uri" locationName:"ruleName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisableTopicRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableTopicRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableTopicRuleInput"}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}
	if s.RuleName != nil && len(*s.RuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisableTopicRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisableTopicRuleOutput) String() string {
	return awsutil.Prettify(s)
}