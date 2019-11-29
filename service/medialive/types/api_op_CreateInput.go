// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/medialive/enums"
)

type CreateInputInput struct {
	_ struct{} `type:"structure"`

	Destinations []InputDestinationRequest `locationName:"destinations" type:"list"`

	InputSecurityGroups []string `locationName:"inputSecurityGroups" type:"list"`

	MediaConnectFlows []MediaConnectFlowRequest `locationName:"mediaConnectFlows" type:"list"`

	Name *string `locationName:"name" type:"string"`

	RequestId *string `locationName:"requestId" type:"string" idempotencyToken:"true"`

	RoleArn *string `locationName:"roleArn" type:"string"`

	Sources []InputSourceRequest `locationName:"sources" type:"list"`

	Tags map[string]string `locationName:"tags" type:"map"`

	Type enums.InputType `locationName:"type" type:"string" enum:"true"`

	// Settings for a private VPC Input.When this property is specified, the input
	// destination addresses will be created in a VPC rather than with public Internet
	// addresses.This property requires setting the roleArn property on Input creation.Not
	// compatible with the inputSecurityGroups property.
	Vpc *InputVpcRequest `locationName:"vpc" type:"structure"`
}

// String returns the string representation
func (s CreateInputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateInputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateInputInput"}
	if s.Vpc != nil {
		if err := s.Vpc.Validate(); err != nil {
			invalidParams.AddNested("Vpc", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateInputOutput struct {
	_ struct{} `type:"structure"`

	Input *Input `locationName:"input" type:"structure"`
}

// String returns the string representation
func (s CreateInputOutput) String() string {
	return awsutil.Prettify(s)
}