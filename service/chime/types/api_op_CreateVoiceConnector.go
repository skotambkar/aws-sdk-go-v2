// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/chime/enums"
)

type CreateVoiceConnectorInput struct {
	_ struct{} `type:"structure"`

	// The AWS Region in which the Amazon Chime Voice Connector is created. Default
	// value: us-east-1.
	AwsRegion enums.VoiceConnectorAwsRegion `type:"string" enum:"true"`

	// The name of the Amazon Chime Voice Connector.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// When enabled, requires encryption for the Amazon Chime Voice Connector.
	//
	// RequireEncryption is a required field
	RequireEncryption *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s CreateVoiceConnectorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVoiceConnectorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVoiceConnectorInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.RequireEncryption == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequireEncryption"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateVoiceConnectorOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector details.
	VoiceConnector *VoiceConnector `type:"structure"`
}

// String returns the string representation
func (s CreateVoiceConnectorOutput) String() string {
	return awsutil.Prettify(s)
}