// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutBucketWebsiteInput struct {
	_ struct{} `type:"structure" payload:"WebsiteConfiguration"`

	// The bucket name.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Container for the request.
	//
	// WebsiteConfiguration is a required field
	WebsiteConfiguration *WebsiteConfiguration `locationName:"WebsiteConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketWebsiteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketWebsiteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketWebsiteInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.WebsiteConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebsiteConfiguration"))
	}
	if s.WebsiteConfiguration != nil {
		if err := s.WebsiteConfiguration.Validate(); err != nil {
			invalidParams.AddNested("WebsiteConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketWebsiteInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type PutBucketWebsiteOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketWebsiteOutput) String() string {
	return awsutil.Prettify(s)
}
