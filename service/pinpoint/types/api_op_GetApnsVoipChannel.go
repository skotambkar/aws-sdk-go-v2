// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetApnsVoipChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetApnsVoipChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApnsVoipChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApnsVoipChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetApnsVoipChannelOutput struct {
	_ struct{} `type:"structure" payload:"APNSVoipChannelResponse"`

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) VoIP channel for an application.
	//
	// APNSVoipChannelResponse is a required field
	APNSVoipChannelResponse *APNSVoipChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetApnsVoipChannelOutput) String() string {
	return awsutil.Prettify(s)
}
