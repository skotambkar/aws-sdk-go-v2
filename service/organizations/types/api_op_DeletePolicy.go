// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeletePolicyInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the policy that you want to delete. You can
	// get the ID from the ListPolicies or ListPoliciesForTarget operations.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a policy ID string
	// requires "p-" followed by from 8 to 128 lower-case letters or digits.
	//
	// PolicyId is a required field
	PolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePolicyInput"}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeletePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePolicyOutput) String() string {
	return awsutil.Prettify(s)
}