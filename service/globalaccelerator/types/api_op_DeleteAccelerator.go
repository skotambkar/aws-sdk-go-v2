// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteAcceleratorInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of an accelerator.
	//
	// AcceleratorArn is a required field
	AcceleratorArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAcceleratorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAcceleratorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAcceleratorInput"}

	if s.AcceleratorArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AcceleratorArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAcceleratorOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAcceleratorOutput) String() string {
	return awsutil.Prettify(s)
}
