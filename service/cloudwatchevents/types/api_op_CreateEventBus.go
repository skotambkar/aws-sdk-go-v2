// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateEventBusInput struct {
	_ struct{} `type:"structure"`

	// If you're creating a partner event bus, this specifies the partner event
	// source that the new event bus will be matched with.
	EventSourceName *string `min:"1" type:"string"`

	// The name of the new event bus.
	//
	// The names of custom event buses can't contain the / character. You can't
	// use the name default for a custom event bus because this name is already
	// used for your account's default event bus.
	//
	// If this is a partner event bus, the name must exactly match the name of the
	// partner event source that this event bus is matched to. This name will include
	// the / character.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateEventBusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEventBusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateEventBusInput"}
	if s.EventSourceName != nil && len(*s.EventSourceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventSourceName", 1))
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

type CreateEventBusOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the new event bus.
	EventBusArn *string `type:"string"`
}

// String returns the string representation
func (s CreateEventBusOutput) String() string {
	return awsutil.Prettify(s)
}