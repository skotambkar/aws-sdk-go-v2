// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAttackInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) for the attack that to be described.
	//
	// AttackId is a required field
	AttackId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAttackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAttackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAttackInput"}

	if s.AttackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttackId"))
	}
	if s.AttackId != nil && len(*s.AttackId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AttackId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeAttackOutput struct {
	_ struct{} `type:"structure"`

	// The attack that is described.
	Attack *AttackDetail `type:"structure"`
}

// String returns the string representation
func (s DescribeAttackOutput) String() string {
	return awsutil.Prettify(s)
}