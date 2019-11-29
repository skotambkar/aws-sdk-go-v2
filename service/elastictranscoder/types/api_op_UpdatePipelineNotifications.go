// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The UpdatePipelineNotificationsRequest structure.
type UpdatePipelineNotificationsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the pipeline for which you want to change notification
	// settings.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic
	// that you want to notify to report job status.
	//
	// To receive notifications, you must also subscribe to the new topic in the
	// Amazon SNS console.
	//
	//    * Progressing: The topic ARN for the Amazon Simple Notification Service
	//    (Amazon SNS) topic that you want to notify when Elastic Transcoder has
	//    started to process jobs that are added to this pipeline. This is the ARN
	//    that Amazon SNS returned when you created the topic.
	//
	//    * Complete: The topic ARN for the Amazon SNS topic that you want to notify
	//    when Elastic Transcoder has finished processing a job. This is the ARN
	//    that Amazon SNS returned when you created the topic.
	//
	//    * Warning: The topic ARN for the Amazon SNS topic that you want to notify
	//    when Elastic Transcoder encounters a warning condition. This is the ARN
	//    that Amazon SNS returned when you created the topic.
	//
	//    * Error: The topic ARN for the Amazon SNS topic that you want to notify
	//    when Elastic Transcoder encounters an error condition. This is the ARN
	//    that Amazon SNS returned when you created the topic.
	//
	// Notifications is a required field
	Notifications *Notifications `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdatePipelineNotificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePipelineNotificationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePipelineNotificationsInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.Notifications == nil {
		invalidParams.Add(aws.NewErrParamRequired("Notifications"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The UpdatePipelineNotificationsResponse structure.
type UpdatePipelineNotificationsOutput struct {
	_ struct{} `type:"structure"`

	// A section of the response body that provides information about the pipeline
	// associated with this notification.
	Pipeline *Pipeline `type:"structure"`
}

// String returns the string representation
func (s UpdatePipelineNotificationsOutput) String() string {
	return awsutil.Prettify(s)
}
