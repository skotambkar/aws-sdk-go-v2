// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ses/enums"
)

// Represents a request to specify the Amazon SNS topic to which Amazon SES
// will publish bounce, complaint, or delivery notifications for emails sent
// with that identity as the Source. For information about Amazon SES notifications,
// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications-via-sns.html).
type SetIdentityNotificationTopicInput struct {
	_ struct{} `type:"structure"`

	// The identity (email address or domain) that you want to set the Amazon SNS
	// topic for.
	//
	// You can only specify a verified identity for this parameter.
	//
	// You can specify an identity by using its name or by using its Amazon Resource
	// Name (ARN). The following examples are all valid identities: sender@example.com,
	// example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com.
	//
	// Identity is a required field
	Identity *string `type:"string" required:"true"`

	// The type of notifications that will be published to the specified Amazon
	// SNS topic.
	//
	// NotificationType is a required field
	NotificationType enums.NotificationType `type:"string" required:"true" enum:"true"`

	// The Amazon Resource Name (ARN) of the Amazon SNS topic. If the parameter
	// is omitted from the request or a null value is passed, SnsTopic is cleared
	// and publishing is disabled.
	SnsTopic *string `type:"string"`
}

// String returns the string representation
func (s SetIdentityNotificationTopicInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetIdentityNotificationTopicInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetIdentityNotificationTopicInput"}

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
type SetIdentityNotificationTopicOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetIdentityNotificationTopicOutput) String() string {
	return awsutil.Prettify(s)
}