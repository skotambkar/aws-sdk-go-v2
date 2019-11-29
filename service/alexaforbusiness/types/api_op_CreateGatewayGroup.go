// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateGatewayGroupInput struct {
	_ struct{} `type:"structure"`

	// A unique, user-specified identifier for the request that ensures idempotency.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `min:"10" type:"string" required:"true" idempotencyToken:"true"`

	// The description of the gateway group.
	Description *string `type:"string"`

	// The name of the gateway group.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateGatewayGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGatewayGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGatewayGroupInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 10))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateGatewayGroupOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the created gateway group.
	GatewayGroupArn *string `type:"string"`
}

// String returns the string representation
func (s CreateGatewayGroupOutput) String() string {
	return awsutil.Prettify(s)
}
