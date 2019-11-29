// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetBucketWebsiteInput struct {
	_ struct{} `type:"structure"`

	// The bucket name for which to get the website configuration.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketWebsiteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketWebsiteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketWebsiteInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketWebsiteInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type GetBucketWebsiteOutput struct {
	_ struct{} `type:"structure"`

	// The name of the error document for the website.
	ErrorDocument *ErrorDocument `type:"structure"`

	// The name of the index document for the website.
	IndexDocument *IndexDocument `type:"structure"`

	// Specifies the redirect behavior of all requests to a website endpoint of
	// an Amazon S3 bucket.
	RedirectAllRequestsTo *RedirectAllRequestsTo `type:"structure"`

	// Rules that define when a redirect is applied and the redirect behavior.
	RoutingRules []RoutingRule `locationNameList:"RoutingRule" type:"list"`
}

// String returns the string representation
func (s GetBucketWebsiteOutput) String() string {
	return awsutil.Prettify(s)
}
