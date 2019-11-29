// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetPermissionPolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the RuleGroup for which you want to get
	// the policy.
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPermissionPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPermissionPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPermissionPolicyInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetPermissionPolicyOutput struct {
	_ struct{} `type:"structure"`

	// The IAM policy attached to the specified RuleGroup.
	Policy *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetPermissionPolicyOutput) String() string {
	return awsutil.Prettify(s)
}
