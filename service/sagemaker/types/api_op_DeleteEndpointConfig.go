// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteEndpointConfigInput struct {
	_ struct{} `type:"structure"`

	// The name of the endpoint configuration that you want to delete.
	//
	// EndpointConfigName is a required field
	EndpointConfigName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEndpointConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEndpointConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEndpointConfigInput"}

	if s.EndpointConfigName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointConfigName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteEndpointConfigOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteEndpointConfigOutput) String() string {
	return awsutil.Prettify(s)
}
