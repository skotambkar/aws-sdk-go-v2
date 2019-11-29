// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociatePhoneNumbersWithVoiceConnectorGroupInput struct {
	_ struct{} `type:"structure"`

	// List of phone numbers, in E.164 format.
	E164PhoneNumbers []string `type:"list"`

	// If true, associates the provided phone numbers with the provided Amazon Chime
	// Voice Connector Group and removes any previously existing associations. If
	// false, does not associate any phone numbers that have previously existing
	// associations.
	ForceAssociate *bool `type:"boolean"`

	// The Amazon Chime Voice Connector group ID.
	//
	// VoiceConnectorGroupId is a required field
	VoiceConnectorGroupId *string `location:"uri" locationName:"voiceConnectorGroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociatePhoneNumbersWithVoiceConnectorGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociatePhoneNumbersWithVoiceConnectorGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociatePhoneNumbersWithVoiceConnectorGroupInput"}

	if s.VoiceConnectorGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociatePhoneNumbersWithVoiceConnectorGroupOutput struct {
	_ struct{} `type:"structure"`

	// If the action fails for one or more of the phone numbers in the request,
	// a list of the phone numbers is returned, along with error codes and error
	// messages.
	PhoneNumberErrors []PhoneNumberError `type:"list"`
}

// String returns the string representation
func (s AssociatePhoneNumbersWithVoiceConnectorGroupOutput) String() string {
	return awsutil.Prettify(s)
}