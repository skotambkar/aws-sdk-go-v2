// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetBucketMetricsConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket containing the metrics configuration to retrieve.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The ID used to identify the metrics configuration.
	//
	// Id is a required field
	Id *string `location:"querystring" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketMetricsConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketMetricsConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketMetricsConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketMetricsConfigurationInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type GetBucketMetricsConfigurationOutput struct {
	_ struct{} `type:"structure" payload:"MetricsConfiguration"`

	// Specifies the metrics configuration.
	MetricsConfiguration *MetricsConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetBucketMetricsConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}