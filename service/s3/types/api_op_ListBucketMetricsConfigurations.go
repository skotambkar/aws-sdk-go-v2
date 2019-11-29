// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListBucketMetricsConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket containing the metrics configurations to retrieve.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The marker that is used to continue a metrics configuration listing that
	// has been truncated. Use the NextContinuationToken from a previously truncated
	// list response to continue the listing. The continuation token is an opaque
	// value that Amazon S3 understands.
	ContinuationToken *string `location:"querystring" locationName:"continuation-token" type:"string"`
}

// String returns the string representation
func (s ListBucketMetricsConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBucketMetricsConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBucketMetricsConfigurationsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *ListBucketMetricsConfigurationsInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type ListBucketMetricsConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// The marker that is used as a starting point for this metrics configuration
	// list response. This value is present if it was sent in the request.
	ContinuationToken *string `type:"string"`

	// Indicates whether the returned list of metrics configurations is complete.
	// A value of true indicates that the list is not complete and the NextContinuationToken
	// will be provided for a subsequent request.
	IsTruncated *bool `type:"boolean"`

	// The list of metrics configurations for a bucket.
	MetricsConfigurationList []MetricsConfiguration `locationName:"MetricsConfiguration" type:"list" flattened:"true"`

	// The marker used to continue a metrics configuration listing that has been
	// truncated. Use the NextContinuationToken from a previously truncated list
	// response to continue the listing. The continuation token is an opaque value
	// that Amazon S3 understands.
	NextContinuationToken *string `type:"string"`
}

// String returns the string representation
func (s ListBucketMetricsConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}
