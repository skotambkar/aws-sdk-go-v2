// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeletePermissionPolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the RuleGroup from which you want to delete
	// the policy.
	//
	// The user making the request must be the owner of the RuleGroup.
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePermissionPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePermissionPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePermissionPolicyInput"}

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

type DeletePermissionPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePermissionPolicyOutput) String() string {
	return awsutil.Prettify(s)
}
