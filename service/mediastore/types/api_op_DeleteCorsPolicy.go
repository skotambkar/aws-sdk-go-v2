// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteCorsPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the container to remove the policy from.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCorsPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCorsPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCorsPolicyInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteCorsPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCorsPolicyOutput) String() string {
	return awsutil.Prettify(s)
}
