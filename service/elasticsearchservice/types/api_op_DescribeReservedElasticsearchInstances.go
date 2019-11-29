// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for parameters to DescribeReservedElasticsearchInstances
type DescribeReservedElasticsearchInstancesInput struct {
	_ struct{} `type:"structure"`

	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// NextToken should be sent in case if earlier API call produced result containing
	// NextToken. It is used for pagination.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The reserved instance identifier filter value. Use this parameter to show
	// only the reservation that matches the specified reserved Elasticsearch instance
	// ID.
	ReservedElasticsearchInstanceId *string `location:"querystring" locationName:"reservationId" type:"string"`
}

// String returns the string representation
func (s DescribeReservedElasticsearchInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Container for results from DescribeReservedElasticsearchInstances
type DescribeReservedElasticsearchInstancesOutput struct {
	_ struct{} `type:"structure"`

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string `type:"string"`

	// List of reserved Elasticsearch instances.
	ReservedElasticsearchInstances []ReservedElasticsearchInstance `type:"list"`
}

// String returns the string representation
func (s DescribeReservedElasticsearchInstancesOutput) String() string {
	return awsutil.Prettify(s)
}