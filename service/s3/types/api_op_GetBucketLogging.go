// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetBucketLoggingInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketLoggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketLoggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketLoggingInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketLoggingInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type GetBucketLoggingOutput struct {
	_ struct{} `type:"structure"`

	// Describes where logs are stored and the prefix that Amazon S3 assigns to
	// all log object keys for a bucket. For more information, see PUT Bucket logging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html)
	// in the Amazon Simple Storage Service API Reference.
	LoggingEnabled *LoggingEnabled `type:"structure"`
}

// String returns the string representation
func (s GetBucketLoggingOutput) String() string {
	return awsutil.Prettify(s)
}
