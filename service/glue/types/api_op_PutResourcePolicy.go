// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/glue/enums"
)

type PutResourcePolicyInput struct {
	_ struct{} `type:"structure"`

	// A value of MUST_EXIST is used to update a policy. A value of NOT_EXIST is
	// used to create a new policy. If a value of NONE or a null value is used,
	// the call will not depend on the existence of a policy.
	PolicyExistsCondition enums.ExistCondition `type:"string" enum:"true"`

	// The hash value returned when the previous policy was set using PutResourcePolicy.
	// Its purpose is to prevent concurrent modifications of a policy. Do not use
	// this parameter if no previous policy has been set.
	PolicyHashCondition *string `min:"1" type:"string"`

	// Contains the policy document to set, in JSON format.
	//
	// PolicyInJson is a required field
	PolicyInJson *string `min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PutResourcePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutResourcePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutResourcePolicyInput"}
	if s.PolicyHashCondition != nil && len(*s.PolicyHashCondition) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyHashCondition", 1))
	}

	if s.PolicyInJson == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyInJson"))
	}
	if s.PolicyInJson != nil && len(*s.PolicyInJson) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyInJson", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutResourcePolicyOutput struct {
	_ struct{} `type:"structure"`

	// A hash of the policy that has just been set. This must be included in a subsequent
	// call that overwrites or updates this policy.
	PolicyHash *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutResourcePolicyOutput) String() string {
	return awsutil.Prettify(s)
}
