// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartOutboundVoiceContactInput struct {
	_ struct{} `type:"structure"`

	// A custom key-value pair using an attribute map. The attributes are standard
	// Amazon Connect attributes, and can be accessed in contact flows just like
	// any other contact attributes.
	//
	// There can be up to 32,768 UTF-8 bytes across all key-value pairs per contact.
	// Attribute keys can include only alphanumeric, dash, and underscore characters.
	Attributes map[string]string `type:"map"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. The token is valid for 7 days after creation. If a contact
	// is already started, the contact ID is returned. If the contact is disconnected,
	// a new contact is started.
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// The identifier of the contact flow for the outbound call.
	//
	// ContactFlowId is a required field
	ContactFlowId *string `type:"string" required:"true"`

	// The phone number of the customer, in E.164 format.
	//
	// DestinationPhoneNumber is a required field
	DestinationPhoneNumber *string `type:"string" required:"true"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `min:"1" type:"string" required:"true"`

	// The queue for the call. If you specify a queue, the phone displayed for caller
	// ID is the phone number specified in the queue. If you do not specify a queue,
	// the queue defined in the contact flow is used. If you do not specify a queue,
	// you must specify a source phone number.
	QueueId *string `type:"string"`

	// The phone number associated with the Amazon Connect instance, in E.164 format.
	// If you do not specify a source phone number, you must specify a queue.
	SourcePhoneNumber *string `type:"string"`
}

// String returns the string representation
func (s StartOutboundVoiceContactInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartOutboundVoiceContactInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartOutboundVoiceContactInput"}

	if s.ContactFlowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContactFlowId"))
	}

	if s.DestinationPhoneNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationPhoneNumber"))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartOutboundVoiceContactOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of this contact within the Amazon Connect instance.
	ContactId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartOutboundVoiceContactOutput) String() string {
	return awsutil.Prettify(s)
}
