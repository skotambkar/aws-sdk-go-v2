// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutNotificationConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// The type of event that causes the notification to be sent. For more information
	// about notification types supported by Amazon EC2 Auto Scaling, see DescribeAutoScalingNotificationTypes.
	//
	// NotificationTypes is a required field
	NotificationTypes []string `type:"list" required:"true"`

	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service
	// (Amazon SNS) topic.
	//
	// TopicARN is a required field
	TopicARN *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutNotificationConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutNotificationConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutNotificationConfigurationInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if s.NotificationTypes == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotificationTypes"))
	}

	if s.TopicARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("TopicARN"))
	}
	if s.TopicARN != nil && len(*s.TopicARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TopicARN", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutNotificationConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutNotificationConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}
