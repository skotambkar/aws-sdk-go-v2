// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeStackSetInput struct {
	_ struct{} `type:"structure"`

	// The name or unique ID of the stack set whose description you want.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStackSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStackSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStackSetInput"}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeStackSetOutput struct {
	_ struct{} `type:"structure"`

	// The specified stack set.
	StackSet *StackSet `type:"structure"`
}

// String returns the string representation
func (s DescribeStackSetOutput) String() string {
	return awsutil.Prettify(s)
}