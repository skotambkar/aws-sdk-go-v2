// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteActivationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the activation that you want to delete.
	//
	// ActivationId is a required field
	ActivationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteActivationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteActivationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteActivationInput"}

	if s.ActivationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActivationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteActivationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteActivationOutput) String() string {
	return awsutil.Prettify(s)
}
