// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetChannelsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetChannelsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetChannelsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetChannelsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetChannelsOutput struct {
	_ struct{} `type:"structure" payload:"ChannelsResponse"`

	// Provides information about the general settings and status of all channels
	// for an application, including channels that aren't enabled for the application.
	//
	// ChannelsResponse is a required field
	ChannelsResponse *ChannelsResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetChannelsOutput) String() string {
	return awsutil.Prettify(s)
}