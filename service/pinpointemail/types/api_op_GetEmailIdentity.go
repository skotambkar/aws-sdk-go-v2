// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/enums"
)

// A request to return details about an email identity.
type GetEmailIdentityInput struct {
	_ struct{} `type:"structure"`

	// The email identity that you want to retrieve details for.
	//
	// EmailIdentity is a required field
	EmailIdentity *string `location:"uri" locationName:"EmailIdentity" type:"string" required:"true"`
}

// String returns the string representation
func (s GetEmailIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEmailIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEmailIdentityInput"}

	if s.EmailIdentity == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailIdentity"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Details about an email identity.
type GetEmailIdentityOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains information about the DKIM attributes for the identity.
	// This object includes the tokens that you use to create the CNAME records
	// that are required to complete the DKIM verification process.
	DkimAttributes *DkimAttributes `type:"structure"`

	// The feedback forwarding configuration for the identity.
	//
	// If the value is true, Amazon Pinpoint sends you email notifications when
	// bounce or complaint events occur. Amazon Pinpoint sends this notification
	// to the address that you specified in the Return-Path header of the original
	// email.
	//
	// When you set this value to false, Amazon Pinpoint sends notifications through
	// other mechanisms, such as by notifying an Amazon SNS topic or another event
	// destination. You're required to have a method of tracking bounces and complaints.
	// If you haven't set up another mechanism for receiving bounce or complaint
	// notifications, Amazon Pinpoint sends an email notification when these events
	// occur (even if this setting is disabled).
	FeedbackForwardingStatus *bool `type:"boolean"`

	// The email identity type.
	IdentityType enums.IdentityType `type:"string" enum:"true"`

	// An object that contains information about the Mail-From attributes for the
	// email identity.
	MailFromAttributes *MailFromAttributes `type:"structure"`

	// An array of objects that define the tags (keys and values) that are associated
	// with the email identity.
	Tags []Tag `type:"list"`

	// Specifies whether or not the identity is verified. In Amazon Pinpoint, you
	// can only send email from verified email addresses or domains. For more information
	// about verifying identities, see the Amazon Pinpoint User Guide (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html).
	VerifiedForSendingStatus *bool `type:"boolean"`
}

// String returns the string representation
func (s GetEmailIdentityOutput) String() string {
	return awsutil.Prettify(s)
}