// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ses/enums"
)

// Represents a request to set whether Amazon SES includes the original email
// headers in the Amazon SNS notifications of a specified type. For information
// about notifications, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications-via-sns.html).
type SetIdentityHeadersInNotificationsEnabledInput struct {
	_ struct{} `type:"structure"`

	// Sets whether Amazon SES includes the original email headers in Amazon SNS
	// notifications of the specified notification type. A value of true specifies
	// that Amazon SES will include headers in notifications, and a value of false
	// specifies that Amazon SES will not include headers in notifications.
	//
	// This value can only be set when NotificationType is already set to use a
	// particular Amazon SNS topic.
	//
	// Enabled is a required field
	Enabled *bool `type:"boolean" required:"true"`

	// The identity for which to enable or disable headers in notifications. Examples:
	// user@example.com, example.com.
	//
	// Identity is a required field
	Identity *string `type:"string" required:"true"`

	// The notification type for which to enable or disable headers in notifications.
	//
	// NotificationType is a required field
	NotificationType enums.NotificationType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s SetIdentityHeadersInNotificationsEnabledInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetIdentityHeadersInNotificationsEnabledInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetIdentityHeadersInNotificationsEnabledInput"}

	if s.Enabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("Enabled"))
	}

	if s.Identity == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identity"))
	}
	if len(s.NotificationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("NotificationType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
type SetIdentityHeadersInNotificationsEnabledOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetIdentityHeadersInNotificationsEnabledOutput) String() string {
	return awsutil.Prettify(s)
}