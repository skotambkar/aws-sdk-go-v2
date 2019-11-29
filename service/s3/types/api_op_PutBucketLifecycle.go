// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutBucketLifecycleInput struct {
	_ struct{} `type:"structure" payload:"LifecycleConfiguration"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Container for lifecycle rules. You can add as many as 1000 rules.
	LifecycleConfiguration *LifecycleConfiguration `locationName:"LifecycleConfiguration" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketLifecycleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketLifecycleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketLifecycleInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}
	if s.LifecycleConfiguration != nil {
		if err := s.LifecycleConfiguration.Validate(); err != nil {
			invalidParams.AddNested("LifecycleConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketLifecycleInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type PutBucketLifecycleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketLifecycleOutput) String() string {
	return awsutil.Prettify(s)
}
