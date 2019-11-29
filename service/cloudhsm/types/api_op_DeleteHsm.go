// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the DeleteHsm operation.
type DeleteHsmInput struct {
	_ struct{} `locationName:"DeleteHsmRequest" type:"structure"`

	// The ARN of the HSM to delete.
	//
	// HsmArn is a required field
	HsmArn *string `locationName:"HsmArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteHsmInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteHsmInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteHsmInput"}

	if s.HsmArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of the DeleteHsm operation.
type DeleteHsmOutput struct {
	_ struct{} `type:"structure"`

	// The status of the operation.
	//
	// Status is a required field
	Status *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteHsmOutput) String() string {
	return awsutil.Prettify(s)
}