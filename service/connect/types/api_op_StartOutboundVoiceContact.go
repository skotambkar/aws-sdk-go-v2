// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartOutboundVoiceContactInput struct {
	_ struct{} `type:"structure"`

	// Specify a custom key-value pair using an attribute map. The attributes are
	// standard Amazon Connect attributes, and can be accessed in contact flows
	// just like any other contact attributes.
	//
	// There can be up to 32,768 UTF-8 bytes across all key-value pairs per contact.
	// Attribute keys can include only alphanumeric, dash, and underscore characters.
	//
	// For example, if you want play a greeting when the customer answers the call,
	// you can pass the customer name in attributes similar to the following:
	Attributes map[string]string `type:"map"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. The token is valid for 7 days after creation. If a contact
	// is already started, the contact ID is returned. If the contact is disconnected,
	// a new contact is started.
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// The identifier for the contact flow to connect the outbound call to.
	//
	// To find the ContactFlowId, open the contact flow you want to use in the Amazon
	// Connect contact flow editor. The ID for the contact flow is displayed in
	// the address bar as part of the URL. For example, the contact flow ID is the
	// set of characters at the end of the URL, after 'contact-flow/' such as 78ea8fd5-2659-4f2b-b528-699760ccfc1b.
	//
	// ContactFlowId is a required field
	ContactFlowId *string `type:"string" required:"true"`

	// The phone number of the customer in E.164 format.
	//
	// DestinationPhoneNumber is a required field
	DestinationPhoneNumber *string `type:"string" required:"true"`

	// The identifier for your Amazon Connect instance. To find the ID of your instance,
	// open the AWS console and select Amazon Connect. Select the alias of the instance
	// in the Instance alias column. The instance ID is displayed in the Overview
	// section of your instance settings. For example, the instance ID is the set
	// of characters at the end of the instance ARN, after instance/, such as 10a4c4eb-f57e-4d4c-b602-bf39176ced07.
	//
	// InstanceId is a required field
	InstanceId *string `min:"1" type:"string" required:"true"`

	// The queue to add the call to. If you specify a queue, the phone displayed
	// for caller ID is the phone number specified in the queue. If you do not specify
	// a queue, the queue used will be the queue defined in the contact flow.
	//
	// To find the QueueId, open the queue you want to use in the Amazon Connect
	// Queue editor. The ID for the queue is displayed in the address bar as part
	// of the URL. For example, the queue ID is the set of characters at the end
	// of the URL, after 'queue/' such as queue/aeg40574-2d01-51c3-73d6-bf8624d2168c.
	QueueId *string `type:"string"`

	// The phone number, in E.164 format, associated with your Amazon Connect instance
	// to use for the outbound call.
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

	// The unique identifier of this contact within your Amazon Connect instance.
	ContactId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartOutboundVoiceContactOutput) String() string {
	return awsutil.Prettify(s)
}
