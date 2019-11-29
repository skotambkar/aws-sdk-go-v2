// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeCacheInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCacheInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCacheInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCacheInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeCacheOutput struct {
	_ struct{} `type:"structure"`

	// The amount of cache in bytes allocated to the a gateway.
	CacheAllocatedInBytes *int64 `type:"long"`

	// The file share's contribution to the overall percentage of the gateway's
	// cache that has not been persisted to AWS. The sample is taken at the end
	// of the reporting period.
	CacheDirtyPercentage *float64 `type:"double"`

	// Percent of application read operations from the file shares that are served
	// from cache. The sample is taken at the end of the reporting period.
	CacheHitPercentage *float64 `type:"double"`

	// Percent of application read operations from the file shares that are not
	// served from cache. The sample is taken at the end of the reporting period.
	CacheMissPercentage *float64 `type:"double"`

	// Percent use of the gateway's cache storage. This metric applies only to the
	// gateway-cached volume setup. The sample is taken at the end of the reporting
	// period.
	CacheUsedPercentage *float64 `type:"double"`

	// An array of strings that identify disks that are to be configured as working
	// storage. Each string have a minimum length of 1 and maximum length of 300.
	// You can get the disk IDs from the ListLocalDisks API.
	DiskIds []string `type:"list"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DescribeCacheOutput) String() string {
	return awsutil.Prettify(s)
}
