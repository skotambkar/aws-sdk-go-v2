// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/medialive/enums"
)

type CreateChannelInput struct {
	_ struct{} `type:"structure"`

	// A standard channel has two encoding pipelines and a single pipeline channel
	// only has one.
	ChannelClass enums.ChannelClass `locationName:"channelClass" type:"string" enum:"true"`

	Destinations []OutputDestination `locationName:"destinations" type:"list"`

	// Encoder Settings
	EncoderSettings *EncoderSettings `locationName:"encoderSettings" type:"structure"`

	InputAttachments []InputAttachment `locationName:"inputAttachments" type:"list"`

	InputSpecification *InputSpecification `locationName:"inputSpecification" type:"structure"`

	// The log level the user wants for their channel.
	LogLevel enums.LogLevel `locationName:"logLevel" type:"string" enum:"true"`

	Name *string `locationName:"name" type:"string"`

	RequestId *string `locationName:"requestId" type:"string" idempotencyToken:"true"`

	Reserved *string `locationName:"reserved" deprecated:"true" type:"string"`

	RoleArn *string `locationName:"roleArn" type:"string"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateChannelInput"}
	if s.Destinations != nil {
		for i, v := range s.Destinations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Destinations", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.EncoderSettings != nil {
		if err := s.EncoderSettings.Validate(); err != nil {
			invalidParams.AddNested("EncoderSettings", err.(aws.ErrInvalidParams))
		}
	}
	if s.InputAttachments != nil {
		for i, v := range s.InputAttachments {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "InputAttachments", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateChannelOutput struct {
	_ struct{} `type:"structure"`

	Channel *Channel `locationName:"channel" type:"structure"`
}

// String returns the string representation
func (s CreateChannelOutput) String() string {
	return awsutil.Prettify(s)
}