// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutBucketLoggingInput struct {
	_ struct{} `type:"structure" payload:"BucketLoggingStatus"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// BucketLoggingStatus is a required field
	BucketLoggingStatus *BucketLoggingStatus `locationName:"BucketLoggingStatus" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketLoggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketLoggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketLoggingInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.BucketLoggingStatus == nil {
		invalidParams.Add(aws.NewErrParamRequired("BucketLoggingStatus"))
	}
	if s.BucketLoggingStatus != nil {
		if err := s.BucketLoggingStatus.Validate(); err != nil {
			invalidParams.AddNested("BucketLoggingStatus", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketLoggingInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type PutBucketLoggingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketLoggingOutput) String() string {
	return awsutil.Prettify(s)
}
