// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutPublicAccessBlockInput struct {
	_ struct{} `type:"structure" payload:"PublicAccessBlockConfiguration"`

	// The name of the Amazon S3 bucket whose PublicAccessBlock configuration you
	// want to set.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The PublicAccessBlock configuration that you want to apply to this Amazon
	// S3 bucket. You can enable the configuration options in any combination. For
	// more information about when Amazon S3 considers a bucket or object public,
	// see The Meaning of "Public" (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// PublicAccessBlockConfiguration is a required field
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `locationName:"PublicAccessBlockConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutPublicAccessBlockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutPublicAccessBlockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutPublicAccessBlockInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.PublicAccessBlockConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublicAccessBlockConfiguration"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutPublicAccessBlockInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type PutPublicAccessBlockOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutPublicAccessBlockOutput) String() string {
	return awsutil.Prettify(s)
}
