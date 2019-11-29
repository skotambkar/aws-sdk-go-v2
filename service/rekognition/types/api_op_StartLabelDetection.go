// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartLabelDetectionInput struct {
	_ struct{} `type:"structure"`

	// Idempotent token used to identify the start request. If you use the same
	// token with multiple StartLabelDetection requests, the same JobId is returned.
	// Use ClientRequestToken to prevent the same job from being accidently started
	// more than once.
	ClientRequestToken *string `min:"1" type:"string"`

	// An identifier you specify that's returned in the completion notification
	// that's published to your Amazon Simple Notification Service topic. For example,
	// you can use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string `min:"1" type:"string"`

	// Specifies the minimum confidence that Amazon Rekognition Video must have
	// in order to return a detected label. Confidence represents how certain Amazon
	// Rekognition is that a label is correctly identified.0 is the lowest confidence.
	// 100 is the highest confidence. Amazon Rekognition Video doesn't return any
	// labels with a confidence level lower than this specified value.
	//
	// If you don't specify MinConfidence, the operation returns labels with confidence
	// values greater than or equal to 50 percent.
	MinConfidence *float64 `type:"float"`

	// The Amazon SNS topic ARN you want Amazon Rekognition Video to publish the
	// completion status of the label detection operation to.
	NotificationChannel *NotificationChannel `type:"structure"`

	// The video in which you want to detect labels. The video must be stored in
	// an Amazon S3 bucket.
	//
	// Video is a required field
	Video *Video `type:"structure" required:"true"`
}

// String returns the string representation
func (s StartLabelDetectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartLabelDetectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartLabelDetectionInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.JobTag != nil && len(*s.JobTag) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobTag", 1))
	}

	if s.Video == nil {
		invalidParams.Add(aws.NewErrParamRequired("Video"))
	}
	if s.NotificationChannel != nil {
		if err := s.NotificationChannel.Validate(); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(aws.ErrInvalidParams))
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			invalidParams.AddNested("Video", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartLabelDetectionOutput struct {
	_ struct{} `type:"structure"`

	// The identifier for the label detection job. Use JobId to identify the job
	// in a subsequent call to GetLabelDetection.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartLabelDetectionOutput) String() string {
	return awsutil.Prettify(s)
}
