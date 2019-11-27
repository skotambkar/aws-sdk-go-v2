// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePushTemplateInput struct {
	_ struct{} `type:"structure" payload:"PushNotificationTemplateRequest"`

	// Specifies the content and settings for a message template that can be used
	// in messages that are sent through a push notification channel.
	//
	// PushNotificationTemplateRequest is a required field
	PushNotificationTemplateRequest *PushNotificationTemplateRequest `type:"structure" required:"true"`

	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"template-name" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePushTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePushTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePushTemplateInput"}

	if s.PushNotificationTemplateRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("PushNotificationTemplateRequest"))
	}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePushTemplateOutput struct {
	_ struct{} `type:"structure" payload:"CreateTemplateMessageBody"`

	// Provides information about an API request or response.
	//
	// CreateTemplateMessageBody is a required field
	CreateTemplateMessageBody *CreateTemplateMessageBody `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreatePushTemplateOutput) String() string {
	return awsutil.Prettify(s)
}
