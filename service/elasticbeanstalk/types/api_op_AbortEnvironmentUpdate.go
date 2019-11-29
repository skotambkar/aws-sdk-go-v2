// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AbortEnvironmentUpdateInput struct {
	_ struct{} `type:"structure"`

	// This specifies the ID of the environment with the in-progress update that
	// you want to cancel.
	EnvironmentId *string `type:"string"`

	// This specifies the name of the environment with the in-progress update that
	// you want to cancel.
	EnvironmentName *string `min:"4" type:"string"`
}

// String returns the string representation
func (s AbortEnvironmentUpdateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AbortEnvironmentUpdateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AbortEnvironmentUpdateInput"}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AbortEnvironmentUpdateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AbortEnvironmentUpdateOutput) String() string {
	return awsutil.Prettify(s)
}
