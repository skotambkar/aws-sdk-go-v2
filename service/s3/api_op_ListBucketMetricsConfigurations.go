// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opListBucketMetricsConfigurations = "ListBucketMetricsConfigurations"

// ListBucketMetricsConfigurationsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Lists the metrics configurations for the bucket. The metrics configurations
// are only for the request metrics of the bucket and do not provide information
// on daily storage metrics. You can have up to 1,000 configurations per bucket.
//
// This operation supports list pagination and does not return more than 100
// configurations at a time. Always check the IsTruncated element in the response.
// If there are no more configurations to list, IsTruncated is set to false.
// If there are more configurations to list, IsTruncated is set to true, and
// there is a value in NextContinuationToken. You use the NextContinuationToken
// value to continue the pagination of the list by passing the value in continuation-token
// in the request to GET the next page.
//
// To use this operation, you must have permissions to perform the s3:GetMetricsConfiguration
// action. The bucket owner has this permission by default. The bucket owner
// can grant this permission to others. For more information about permissions,
// see Permissions Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// For more information about metrics configurations and CloudWatch request
// metrics, see Monitoring Metrics with Amazon CloudWatch (https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html).
//
// The following operations are related to ListBucketMetricsConfigurations:
//
//    * PutBucketMetricsConfiguration
//
//    * GetBucketMetricsConfiguration
//
//    * DeleteBucketMetricsConfiguration
//
//    // Example sending a request using ListBucketMetricsConfigurationsRequest.
//    req := client.ListBucketMetricsConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListBucketMetricsConfigurations
func (c *Client) ListBucketMetricsConfigurationsRequest(input *types.ListBucketMetricsConfigurationsInput) ListBucketMetricsConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListBucketMetricsConfigurations,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?metrics",
	}

	if input == nil {
		input = &types.ListBucketMetricsConfigurationsInput{}
	}

	req := c.newRequest(op, input, &types.ListBucketMetricsConfigurationsOutput{})
	return ListBucketMetricsConfigurationsRequest{Request: req, Input: input, Copy: c.ListBucketMetricsConfigurationsRequest}
}

// ListBucketMetricsConfigurationsRequest is the request type for the
// ListBucketMetricsConfigurations API operation.
type ListBucketMetricsConfigurationsRequest struct {
	*aws.Request
	Input *types.ListBucketMetricsConfigurationsInput
	Copy  func(*types.ListBucketMetricsConfigurationsInput) ListBucketMetricsConfigurationsRequest
}

// Send marshals and sends the ListBucketMetricsConfigurations API request.
func (r ListBucketMetricsConfigurationsRequest) Send(ctx context.Context) (*ListBucketMetricsConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBucketMetricsConfigurationsResponse{
		ListBucketMetricsConfigurationsOutput: r.Request.Data.(*types.ListBucketMetricsConfigurationsOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListBucketMetricsConfigurationsResponse is the response type for the
// ListBucketMetricsConfigurations API operation.
type ListBucketMetricsConfigurationsResponse struct {
	*types.ListBucketMetricsConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBucketMetricsConfigurations request.
func (r *ListBucketMetricsConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
