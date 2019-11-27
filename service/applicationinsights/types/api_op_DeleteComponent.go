// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteComponentInput struct {
	_ struct{} `type:"structure"`

	// The name of the component.
	//
	// ComponentName is a required field
	ComponentName *string `type:"string" required:"true"`

	// The name of the resource group.
	//
	// ResourceGroupName is a required field
	ResourceGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteComponentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteComponentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteComponentInput"}

	if s.ComponentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ComponentName"))
	}

	if s.ResourceGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteComponentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteComponentOutput) String() string {
	return awsutil.Prettify(s)
}
