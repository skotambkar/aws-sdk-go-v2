// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/dlm/enums"
)

type UpdateLifecyclePolicyInput struct {
	_ struct{} `type:"structure"`

	// A description of the lifecycle policy.
	Description *string `type:"string"`

	// The Amazon Resource Name (ARN) of the IAM role used to run the operations
	// specified by the lifecycle policy.
	ExecutionRoleArn *string `type:"string"`

	// The configuration of the lifecycle policy. You cannot update the policy type
	// or the resource type.
	PolicyDetails *PolicyDetails `type:"structure"`

	// The identifier of the lifecycle policy.
	//
	// PolicyId is a required field
	PolicyId *string `location:"uri" locationName:"policyId" type:"string" required:"true"`

	// The desired activation state of the lifecycle policy after creation.
	State enums.SettablePolicyStateValues `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateLifecyclePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateLifecyclePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateLifecyclePolicyInput"}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}
	if s.PolicyDetails != nil {
		if err := s.PolicyDetails.Validate(); err != nil {
			invalidParams.AddNested("PolicyDetails", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateLifecyclePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateLifecyclePolicyOutput) String() string {
	return awsutil.Prettify(s)
}
