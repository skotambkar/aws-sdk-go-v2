// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the CreateHapgRequest action.
type CreateHapgInput struct {
	_ struct{} `type:"structure"`

	// The label of the new high-availability partition group.
	//
	// Label is a required field
	Label *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateHapgInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHapgInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHapgInput"}

	if s.Label == nil {
		invalidParams.Add(aws.NewErrParamRequired("Label"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of the CreateHAPartitionGroup action.
type CreateHapgOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the high-availability partition group.
	HapgArn *string `type:"string"`
}

// String returns the string representation
func (s CreateHapgOutput) String() string {
	return awsutil.Prettify(s)
}
