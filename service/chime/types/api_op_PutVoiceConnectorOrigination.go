// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutVoiceConnectorOriginationInput struct {
	_ struct{} `type:"structure"`

	// The origination setting details to add.
	//
	// Origination is a required field
	Origination *Origination `type:"structure" required:"true"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s PutVoiceConnectorOriginationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutVoiceConnectorOriginationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutVoiceConnectorOriginationInput"}

	if s.Origination == nil {
		invalidParams.Add(aws.NewErrParamRequired("Origination"))
	}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}
	if s.Origination != nil {
		if err := s.Origination.Validate(); err != nil {
			invalidParams.AddNested("Origination", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutVoiceConnectorOriginationOutput struct {
	_ struct{} `type:"structure"`

	// The updated origination setting details.
	Origination *Origination `type:"structure"`
}

// String returns the string representation
func (s PutVoiceConnectorOriginationOutput) String() string {
	return awsutil.Prettify(s)
}
