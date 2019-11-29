// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeregisterInstanceInput struct {
	_ struct{} `type:"structure"`

	// The instance ID.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterInstanceInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterInstanceOutput) String() string {
	return awsutil.Prettify(s)
}