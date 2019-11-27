// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateSmsTemplateInput struct {
	_ struct{} `type:"structure" payload:"SMSTemplateRequest"`

	// Specifies the content and settings for a message template that can be used
	// in text messages that are sent through the SMS channel.
	//
	// SMSTemplateRequest is a required field
	SMSTemplateRequest *SMSTemplateRequest `type:"structure" required:"true"`

	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"template-name" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateSmsTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSmsTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSmsTemplateInput"}

	if s.SMSTemplateRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("SMSTemplateRequest"))
	}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateSmsTemplateOutput struct {
	_ struct{} `type:"structure" payload:"MessageBody"`

	// Provides information about an API request or response.
	//
	// MessageBody is a required field
	MessageBody *MessageBody `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateSmsTemplateOutput) String() string {
	return awsutil.Prettify(s)
}
